# Running Web Tests with Sauce Labs

Before starting the tests you need to export some environmnent variables and start Sauce Connect. You can dowload Sauce Connect at https://wiki.saucelabs.com/display/DOCS/Setting+Up+Sauce+Connect+Proxy

In a separate terminal run the following:
```sh
export SAUCE_USERNAME=<your user name>
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
```

Then to actually run the SauceLabs tests, use the following:
```sh
bazel test --test_tag_filters=sauce --test_output=streamed ...
```
