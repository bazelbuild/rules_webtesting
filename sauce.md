# Running Web Tests with Sauce Labs

There are two Sauce browsers defined and used by tests.

1.  //browsers/sauce:chrome-win10, requires that SauceConnect is already
    running.

2.  //browsers/sauce:chrome-win10-connect, will start an instance of
    SauceConnect for each browser target.

## //browsers/sauce:chrome-win10 tests

Before starting the tests you need to export some environment variables and
start Sauce Connect. You can dowload Sauce Connect at
https://wiki.saucelabs.com/display/DOCS/Setting+Up+Sauce+Connect+Proxy

In a separate terminal run the following:

```sh
export SAUCE_USERNAME=<your username>
export SAUCE_ACCESS_KEY=<your access key>
export TUNNEL_IDENTIFIER=<whatever you want>
sc -i $TUNNEL_IDENTIFIER
```

Wait for the message: "Sauce Connect is up, you may start your tests."

Then in the terminal where you are going to run Bazel do the following:

```sh
export SAUCE_USERNAME=<your user name>
export SAUCE_ACCESS_KEY=<your access key>
export TUNNEL_IDENTIFIER=<whatever you want>
export BUILD_TAG=<whatever you want>
```

Then to run the tests, use the following:

```sh
bazel test --test_tag_filters=chrome-win10 --test_output=streamed ...
```

## //browsers/sauce:chrome-win10-connect tests

To run the tests, use the following:

```sh
bazel test --test_tag_filters=chrome-win10-connect --test_output=streamed ...
```
