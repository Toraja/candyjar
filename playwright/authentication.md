# Authentication

## Tips

### The page might or might not ask for authentication

This should not happen in a typical scenario, but can happen when using an existing browser instance.  
The script must watch whether the page is the one that it tried to open, or an authentication page.  
Playwright does not have a built-in API for this, but you can use `Promise.race` to wait for either of the two URLs to load.

```ts
await page.goto(expectedPageUrl);
await Promise.race([
  page.waitForURL(expectedPageUrl),
  page.waitForURL(signinPageUrl),
]);
if (page.url().includes(signinPage2Url)) {
  // sign in
}

// continue with the rest of the script
```

### Wait for MFA sign in

If sign-in requires MFA and user interaction, wait for the landing page after sign-in.  
Once the user finishes sign-in, the script can continue with the rest of the steps.  
Also it will be nice to tell a user that they need to perform sign-in manually.

```ts
await page.evaluate(() => {
  const bar = document.createElement('div');
  bar.style.cssText = `
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    background: #0055aa;
    color: white;
    padding: 14px;
    font-size: 16px;
    text-align: center;
    z-index: 10000;
    font-family: sans-serif;
  `;
  bar.textContent =
    '🟡 Please finish manual sign in. The rest of the steps will continue automatically.';
  document.body.appendChild(bar);
});
await page.waitForURL(baseUrl);

// continue with the rest of the script
```
