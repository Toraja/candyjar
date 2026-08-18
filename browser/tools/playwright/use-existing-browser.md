# Use Existing Browser

## Prerequisites

### Start a browser instance with the `--remote-debugging-port` flag

#### Edge

The value of `--user-data-dir` can be any directory that doesn't exist or is empty. It is required to avoid conflicts with an existing browser instance.

```cmd
C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe" --remote-debugging-port=9222 --user-data-dir="%TEMP%\edge-cdp"
```

In `package.json` (`npm` uses `cmd` on Windows):
```json
{
  "scripts": {
    "msedge-cdp": "start \"\" \"C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe\" --remote-debugging-port=9222 --user-data-dir=\"%TEMP%\\edge-cdp\"",
  },
}
```

## Playwright Script

### Connect to the browser instance in script

```ts
import { chromium } from 'playwright';

(async () => {
  const browser = await chromium.connectOverCDP('http://localhost:9222');
  const context = browser.contexts()[0];  // Use existing context
  const page = context.pages()[0];        // Use existing page

  // Interact with the existing browser instance
  await page.goto('https://example.com');
  console.log(await page.title());
  // ... other interactions ...
})();
```

## AI Agents

### Setup MCP server

This configuration allows you to connect to an existing browser instance using the Chrome DevTools Protocol (CDP) endpoint.

#### VS Code

`mcp.json`
```json
{
  "servers": {
    "playwright-mcp-cdp": { // arbitrary name for the server
      "type": "stdio",
      "command": "npx",
      "args": [
        "@playwright/mcp@latest",
        "--cdp-endpoint",
        "http://localhost:9222"
      ]
    }
  }
}
```

### Ask the agent to use the existing browser instance

Ask the agent to do things on the existing browser instance. It is better to specify the server name defined in `mcp.json` to ensure it uses the correct configuration.

## Run Playwright Script First Then Have AI Agent Taken Over

Ask the agent to:
1. Start a browser instance with the `--remote-debugging-port` flag if it is not already running.
1. Run the Playwright script.
  - Make sure the script does not close the browser instance after it finishes. You can achieve this by not calling `browser.close()` in the script.
1. Do whatever you want it to do on the browser instance.
