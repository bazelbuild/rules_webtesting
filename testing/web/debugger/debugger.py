# Copyright 2017 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Web Test Launcher Debugger Front-End."""

import code
import getpass
import hashlib
import json
import socket
import sys
import urllib


class Debugger:
  """Debugger connects to the WTL debugger and sends and receives messages."""

  def __init__(self, port, host="localhost"):
    self._conn = socket.create_connection(address=(host, port))
    self._file = self._conn.makefile(mode="w")
    self._decoder = json.JSONDecoder()
    self._next_id = 1
    self._buffer = ""

  def _get_next_id(self):
    id = self._next_id
    self._next_id = self._next_id + 1
    return id

  def _read_next(self):
    while True:
      try:
        self._buffer = self._buffer.strip()
        obj, p = self._decoder.raw_decode(self._buffer)
        self._buffer = self._buffer[p:]
        return obj
      except:
        pass
      s = self._conn.recv(4096).decode("utf-8")
      if s == "":
        quit()
      self._buffer = self._buffer + s

  def _read_until_waiting(self):
    while True:
      n = self._read_next()
      print(n)
      if n["status"] != "running":
        return

  def step(self):
    """Execute the waiting WebDriver command and stop at the next one."""
    id = self._get_next_id()
    json.dump(obj={"id": id, "command": "step"}, fp=self._file)
    self._file.flush()
    self._read_until_waiting()

  def run(self):
    """Execute WebDriver commands until a breakpoint is reached."""
    id = self._get_next_id()
    json.dump(obj={"id": id, "command": "continue"}, fp=self._file)
    self._file.flush()
    self._read_until_waiting()

  def stop(self):
    """Quit WTL and the debugger."""
    id = self._get_next_id()
    json.dump(obj={"id": id, "command": "stop"}, fp=self._file)
    self._file.flush()
    quit()

  def set_breakpoint(self, path=None, methods=None, body=None):
    """Set a WTL breakpoint.

    Args:
      path: string, Go regular expression to compare to WebDriver command paths.
      methods: list of strings, a list of HTTP methods ("POST", "GET", etc).
      body: string, Go regular expression to compare to body of WebDriver
        command.

    Returns:
      int, id of the breakpoint (can be used in delete_breakpoint command).
    """
    id = self._get_next_id()

    bp = {"id": id}
    if path:
      bp["path"] = path
    if methods:
      bp["methods"] = methods
    if body:
      bp["body"] = body

    json.dump(
        obj={
            "id": id,
            "command": "set breakpoint",
            "breakpoint": bp,
        },
        fp=self._file)
    self._file.flush()
    self._read_until_waiting()
    return id

  def delete_breakpoint(self, breakpoint_id):
    """Delete a previously set WTL breakpoint.

    Args:
      breakpoint_id: int, id of the breakpoint to delete.
    """
    id = self._get_next_id()

    json.dump(
        obj={
            "id": id,
            "command": "delete breakpoint",
            "breakpoint": {
                "id": breakpoint_id
            },
        },
        fp=self._file)
    self._file.flush()
    self._read_until_waiting()


def collect_analytics():
  try:
    urllib.urlopen(
        "http://www.google-analytics.com/collect?v=1&aip=1&tid=UA-52159295-3"
        "&t=screenview&cd=start&an=WTL+Debugger&uid=" +
        hashlib.md5(getpass.getuser()).hexdigest()).close
  except:
    # Error collecting usage
    pass


def main(args):
  host = "localhost"
  if len(args) == 2:
    port = args[1]
  elif len(args) == 3:
    host = args[1]
    port = args[2]
  else:
    print("Usage %s [host] port")
    quit()

  wtl = Debugger(host=host, port=port)

  collect_analytics()

  code.interact(
      """
\033[95m\033[1mPython Interactive Console\033[0m

Debugger Commands:
    wtl.run(): Run test until next WTL breakpoint.
    wtl.step(): Execute the current waiting command and break at the next one.
    wtl.stop(): Quit WTL and the Debugger.
    help(wtl): Additional debugger commands.
""",
      local={"wtl": wtl})


if __name__ == "__main__":
  main(sys.argv)
