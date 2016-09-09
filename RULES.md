# Rules reference

<a name="web_test_suite"></a>
## web_test_suite

<pre>
web_test_suite(<a href="#web_test_suite.name">name</a>, <a href="#web_test_suite.test">test</a>, <a href="#web_test_suite.browsers">browsers</a>, <a href="#web_test_suite.data">data</a>, <a href="#web_test_suite.deprecation">deprecation</a>, <a href="#web_test_suite.shard_count">shard_count</a>, <a href="#web_test_suite.size">size</a>, <a href="#web_test_suite.args">args</a>, <a href="#web_test_suite.tags">tags</a>, <a href="#web_test_suite.timeout">timeout</a>, <a href="#web_test_suite.browser_overrides">browser_overrides</a>, <a href="#web_test_suite.flaky">flaky</a>, <a href="#web_test_suite.config">config</a>, <a href="#web_test_suite.visibility">visibility</a>, <a href="#web_test_suite.local">local</a>)
</pre>

Defines a test_suite of web_test targets to be run.


<a name="web_test_suite_args"></a>
### Attributes


<table class="params-table">
  <colgroup>
    <col class="col-param" />
    <col class="col-description" />
  </colgroup>
  <tbody>
    <tr id="web_test_suite.name">
      <td><code>name</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#name">Name</a>; Required</code></p>
        <p>Name; required. A unique name for this rule.</p>
      </td>
    </tr>
    <tr id="web_test_suite.test">
      <td><code>test</code></td>
      <td>
        <p><code>Unknown; Required</code></p>
        <p>Label; required. A single <em>_test or </em>_binary target. The test that
web_test should run with the specified browser.</p>
      </td>
    </tr>
    <tr id="web_test_suite.browsers">
      <td><code>browsers</code></td>
      <td>
        <p><code>Unknown; Required</code></p>
        <p>List of labels; required. The browsers with which to run the test.</p>
      </td>
    </tr>
    <tr id="web_test_suite.data">
      <td><code>data</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>Label List; optional.</p>
      </td>
    </tr>
    <tr id="web_test_suite.deprecation">
      <td><code>deprecation</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>String; optional.</p>
      </td>
    </tr>
    <tr id="web_test_suite.shard_count">
      <td><code>shard_count</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>Integer; optional; default is 1.</p>
      </td>
    </tr>
    <tr id="web_test_suite.size">
      <td><code>size</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>String; optional; default is 'large'</p>
      </td>
    </tr>
    <tr id="web_test_suite.args">
      <td><code>args</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>String list; optional; list of arguments to pass to test.</p>
      </td>
    </tr>
    <tr id="web_test_suite.tags">
      <td><code>tags</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>String list; optional.</p>
      </td>
    </tr>
    <tr id="web_test_suite.timeout">
      <td><code>timeout</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>String; optional.</p>
      </td>
    </tr>
    <tr id="web_test_suite.browser_overrides">
      <td><code>browser_overrides</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>Dictionary; optional; default is an empty dictionary. A
dictionary mapping from browser names to browser-specific web_test
attributes, such as shard_count, flakiness, timeout, etc. For example:
{'\browsers:chrome-native': {'shard_count': 3, 'flaky': 1}
'\browsers:firefox-native': {'shard_count': 1, 'size': 'medium',
'timeout': 100}}.</p>
      </td>
    </tr>
    <tr id="web_test_suite.flaky">
      <td><code>flaky</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>Boolean; optional.</p>
      </td>
    </tr>
    <tr id="web_test_suite.config">
      <td><code>config</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>Label; optional; default is //external:web_test_default_config.
Configuration of web test features.</p>
      </td>
    </tr>
    <tr id="web_test_suite.visibility">
      <td><code>visibility</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>Label List; optional.</p>
      </td>
    </tr>
    <tr id="web_test_suite.local">
      <td><code>local</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
        <p>Boolean; optional.</p>
      </td>
    </tr>
  </tbody>
</table>

## browser

<pre>
browser(<a href="#browser.name">name</a>, <a href="#browser.data">data</a>, <a href="#browser.disabled">disabled</a>, <a href="#browser.metadata">metadata</a>, <a href="#browser.required_tags">required_tags</a>)
</pre>

Defines a browser configuration for use with web_test.


<a name="browser_args"></a>
### Attributes


<table class="params-table">
  <colgroup>
    <col class="col-param" />
    <col class="col-description" />
  </colgroup>
  <tbody>
    <tr id="browser.name">
      <td><code>name</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#name">Name</a>; Required</code></p>
        <p>A unique name for this rule.</p>
      </td>
    </tr>
    <tr id="browser.data">
      <td><code>data</code></td>
      <td>
        <p><code>List of <a href="http://bazel.io/docs/build-ref.html#labels">labels</a>; Optional</code></p>
        <p>Runtime dependencies needed for this browser.</p>
      </td>
    </tr>
    <tr id="browser.disabled">
      <td><code>disabled</code></td>
      <td>
        <p><code>String; Optional</code></p>
        <p>If set, then a no-op test will be run for all tests using
this browser.</p>
      </td>
    </tr>
    <tr id="browser.metadata">
      <td><code>metadata</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#labels">Label</a>; Required</code></p>
        <p>The web_test metadata file that defines how this browser
is launched and default capabilities for this browser.</p>
      </td>
    </tr>
    <tr id="browser.required_tags">
      <td><code>required_tags</code></td>
      <td>
        <p><code>List of strings; Optional</code></p>
        <p>A list of tags that all web_tests using this browser
should have. Examples include "requires-network", "local", etc.</p>
      </td>
    </tr>
  </tbody>
</table>

## web_test

<pre>
web_test(<a href="#web_test.name">name</a>, <a href="#web_test.data">data</a>, <a href="#web_test.browser">browser</a>, <a href="#web_test.config">config</a>, <a href="#web_test.test">test</a>)
</pre>

Runs a provided test against a provided browser configuration.


<a name="web_test_args"></a>
### Attributes


<table class="params-table">
  <colgroup>
    <col class="col-param" />
    <col class="col-description" />
  </colgroup>
  <tbody>
    <tr id="web_test.name">
      <td><code>name</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#name">Name</a>; Required</code></p>
        <p>A unique name for this rule.</p>
      </td>
    </tr>
    <tr id="web_test.data">
      <td><code>data</code></td>
      <td>
        <p><code>List of <a href="http://bazel.io/docs/build-ref.html#labels">labels</a>; Optional</code></p>
        <p>Additional runtime dependencies for the test.</p>
      </td>
    </tr>
    <tr id="web_test.browser">
      <td><code>browser</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#labels">Label</a>; Required</code></p>
        <p>A browser configuration that defines the type of browser used for
this test.</p>
      </td>
    </tr>
    <tr id="web_test.config">
      <td><code>config</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#labels">Label</a>; Optional</code></p>
        <p>Additional configuration that overrides the configuration in browser.</p>
      </td>
    </tr>
    <tr id="web_test.test">
      <td><code>test</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#labels">Label</a>; Required</code></p>
        <p>The test that will be run against the provided browser.</p>
      </td>
    </tr>
  </tbody>
</table>

## web_test_config

<pre>
web_test_config(<a href="#web_test_config.name">name</a>, <a href="#web_test_config.data">data</a>, <a href="#web_test_config.configs">configs</a>, <a href="#web_test_config.metadata">metadata</a>)
</pre>

A browser-independent configuration that can be used across multiple web_tests.


<a name="web_test_config_args"></a>
### Attributes


<table class="params-table">
  <colgroup>
    <col class="col-param" />
    <col class="col-description" />
  </colgroup>
  <tbody>
    <tr id="web_test_config.name">
      <td><code>name</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#name">Name</a>; Required</code></p>
        <p>A unique name for this rule.</p>
      </td>
    </tr>
    <tr id="web_test_config.data">
      <td><code>data</code></td>
      <td>
        <p><code>List of <a href="http://bazel.io/docs/build-ref.html#labels">labels</a>; Optional</code></p>
        <p>Additional files that this web_test_config depends on at runtime.</p>
      </td>
    </tr>
    <tr id="web_test_config.configs">
      <td><code>configs</code></td>
      <td>
        <p><code>List of <a href="http://bazel.io/docs/build-ref.html#labels">labels</a>; Optional</code></p>
        <p>A list of web_test_config rules that this rule inherits from.
Configuration in rules later in the list will override configuration
earlier in the list.</p>
      </td>
    </tr>
    <tr id="web_test_config.metadata">
      <td><code>metadata</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#labels">Label</a>; Optional</code></p>
        <p>A web_test metadata file with configuration that will override
all other configuration.</p>
      </td>
    </tr>
  </tbody>
</table>

## web_test_named_executable

<pre>
web_test_named_executable(<a href="#web_test_named_executable.name">name</a>, <a href="#web_test_named_executable.data">data</a>, <a href="#web_test_named_executable.alt_name">alt_name</a>, <a href="#web_test_named_executable.executable">executable</a>)
</pre>

Defines a executable that can be located by name.


<a name="web_test_named_executable_args"></a>
### Attributes


<table class="params-table">
  <colgroup>
    <col class="col-param" />
    <col class="col-description" />
  </colgroup>
  <tbody>
    <tr id="web_test_named_executable.name">
      <td><code>name</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#name">Name</a>; Required</code></p>
        <p>A unique name for this rule.</p>
      </td>
    </tr>
    <tr id="web_test_named_executable.data">
      <td><code>data</code></td>
      <td>
        <p><code>List of <a href="http://bazel.io/docs/build-ref.html#labels">labels</a>; Optional</code></p>
        <p>Runtime dependencies for the executable.</p>
      </td>
    </tr>
    <tr id="web_test_named_executable.alt_name">
      <td><code>alt_name</code></td>
      <td>
        <p><code>String; Optional</code></p>
        <p>If supplied, is used instead of name to lookup the executable.</p>
      </td>
    </tr>
    <tr id="web_test_named_executable.executable">
      <td><code>executable</code></td>
      <td>
        <p><code><a href="http://bazel.io/docs/build-ref.html#labels">Label</a>; Required</code></p>
        <p>The executable that will be returned for name or alt_name.</p>
      </td>
    </tr>
  </tbody>
</table>

<a name="web_test_repositories"></a>
## web_test_repositories

<pre>
web_test_repositories(<a href="#web_test_repositories.prefix">prefix</a>, <a href="#web_test_repositories.java">java</a>, <a href="#web_test_repositories.go">go</a>, <a href="#web_test_repositories.launcher">launcher</a>, <a href="#web_test_repositories.merger">merger</a>, <a href="#web_test_repositories.default_config">default_config</a>)
</pre>




<a name="web_test_repositories_args"></a>
### Attributes


<table class="params-table">
  <colgroup>
    <col class="col-param" />
    <col class="col-description" />
  </colgroup>
  <tbody>
    <tr id="web_test_repositories.prefix">
      <td><code>prefix</code></td>
      <td>
        <p><code>String; Optional</code></p>
      </td>
    </tr>
    <tr id="web_test_repositories.java">
      <td><code>java</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
      </td>
    </tr>
    <tr id="web_test_repositories.go">
      <td><code>go</code></td>
      <td>
        <p><code>Unknown; Optional</code></p>
      </td>
    </tr>
    <tr id="web_test_repositories.launcher">
      <td><code>launcher</code></td>
      <td>
        <p><code>String; Optional</code></p>
      </td>
    </tr>
    <tr id="web_test_repositories.merger">
      <td><code>merger</code></td>
      <td>
        <p><code>String; Optional</code></p>
      </td>
    </tr>
    <tr id="web_test_repositories.default_config">
      <td><code>default_config</code></td>
      <td>
        <p><code>String; Optional</code></p>
      </td>
    </tr>
  </tbody>
</table>
