> #### ⚠️ This is currently an alpha version
>
> This is not supported by the k6 team, and is worked on by a(n) individual contributor(s).
> It may also break in the future as both xk6 and playwright-go evolve.
> Any issues with the tool should be raised [here](https://github.com/nicholasvuono/xk6-playwright/issues).
>
>It is not production ready yet, but definitely works as intended! Please enjoy the tool! 

<br><br>
<div align="center">
   <img src="images/xk6_logo.PNG" width="400" alt="pdq"/><br>
   <h1><b>xk6 playwright</b></h1><br>
   <p>k6 extension that adds support for browser automation and end-to-end web testing using <a href="https://github.com/mxschmitt/playwright-go" target="_blank">playwright-go</a></p>
   <p>Special thanks to all the contributors over at <a href="https://github.com/grafana/k6/graphs/contributors" target="_blank">k6</a> and <a href="https://github.com/mxschmitt/playwright-go/graphs/contributors" target="_blank">playwright-go</a>
   <p>Here's to open source!</p>

   <a href="https://github.com/nicholasvuono/xk6-playwright/releases"><img alt="GitHub license" src="https://img.shields.io/badge/release-v0.1.1-blue"></a>
   <a href="https://goreportcard.com/report/github.com/wosp-io/xk6-playwright"><img src="https://goreportcard.com/badge/github.com/wosp-io/xk6-playwright?dummy=unused" alt="Go Report Card"></a>
   <a href="https://github.com/wosp-io/xk6-playwright/blob/main/LICENSE"><img alt="GitHub license" src="https://img.shields.io/github/license/wosp-io/xk6-playwright?color=red"></a>
</div>

----
This project was inspired by <a href="https://github.com/grafana/xk6-browser" target="_blank">xk6-browser</a>. Having seen the release we were excited to play around with the tool, but while using it we ran into some issues around context, page navigation, typing and button clicks. Having previously worked with <a href="https://github.com/mxschmitt/playwright-go" target="_blank">playwright-go</a> we thought it would be a great idea to create an extension around this so we had something we know would work to our liking. Thus <a href="https://github.com/wosp-io/xk6-playwright" target="_blank">xk6 playwright</a> was born!

NOTE: we totally understand at the time of writing this that xk6-browser is not yet production ready, and that it is currently an early beta that has been released to the public that will evolve and get better over time. However, we still saw validity in creating this extension aiming to support something we have used and know works to our liking. Competition is not intended, we just want to make cool things that help us do our jobs!

</br>

## Build from source

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git
- Clone the k6 repository (NOTE: make sure you are in the k6 repo base directory before attempting the below steps)

Then:

1. Install `xk6`:
  ```shell
  go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  xk6 build --output xk6-playwright --with github.com/wosp-io/xk6-playwright
  ```

  This will create a `xk6-playwright` binary file in the current working directory. This file can be used exactly the same as the main `k6` binary, with the addition of being able to run xk6-browser scripts.

3. Run scripts that import `k6/x/playwright` with the new `xk6-playwright` binary. On Linux and macOS make sure this is done by referencing the file in the current directory, e.g. `./xk6-playwright run <script>`, or you can place it somewhere in your `PATH` so that it can be run from anywhere on your system.

</br>

## Simplest Working Example

```JavaScript
import pw from 'k6/x/playwright';

export default function () {
  pw.launch()
  pw.newPage()
  pw.goto("https://www.google.com/", {waitUntil: 'networkidle'})
  pw.waitForSelector("//html/body/div[1]/div[2]", {state: 'visible'})
  pw.kill()
}
```

</br>

## Currently Supported Actions

[Playwright API](https://playwright.dev/docs/api/class-playwright) coverage is as follows:

| Action | Encompassed Playwright Function(s) | Description |
|   :---   | :--- | :--- |
| launch() | [`Run()`](https://pkg.go.dev/github.com/mxschmitt/playwright-go#Run) & [`Launch()`](https://pkg.go.dev/github.com/mxschmitt/playwright-go#BrowserType.Launch) | starts playwright client and launches browser|
| newPage() | [`NewPage()`](https://pkg.go.dev/github.com/mxschmitt/playwright-go#Browser.NewPage) | opens up a new page within the browser |
| waitForSelector() | [`WaitForSelector()`](https://pkg.go.dev/github.com/mxschmitt/playwright-go#Page.WaitForSelector) | waits for an element to be on the page based on the provided selector |
| click() | [`Click()`](https://pkg.go.dev/github.com/mxschmitt/playwright-go#Page.Click) | clicks an element on the page based on the provided selector |
| type() | [`Type()`](https://pkg.go.dev/github.com/mxschmitt/playwright-go#Page.Type) | types in an 'input' element on the page based on the provided selector and string to be entered|

NOTE: the above 'Encompassed Playwright Function(s)' will link to the [playwright-go package documentation](https://pkg.go.dev/github.com/mxschmitt/playwright-go#section-readme) to give an in-depth overview of how these functions will behave from a low-level perspective. 

If you would like a high-level perspective on how these actions work you will be better served with the [Playwright API Documentation](https://playwright.dev/docs/api/class-playwright) 

</br>

## Contributing

1. Fork it (<https://github.com/your-github-user/xk6-playwright/fork>)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

</br>

## Contributors

[Nick Vuono](https://github.com/nicholasvuono) - creator and maintainer
