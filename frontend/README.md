# Getting Started with ViteJS & Typescript React

This project was bootstrapped with many ViteJS examples:

- https://github.com/MR-Mostafa/Viact
- https://github.com/suren-atoyan/react-pwa

## Available Scripts

In the project directory, you can run:

### `yarn start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `yarn test`

Launches the test runner.

### `yarn build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

### Other commands

Other commands can be found in `package.json` file.

## Learn More

To learn React, check out the [React documentation](https://reactjs.org/).

## Debug tools

### Why did you render ?

That library will display the reasons of a component re-render.

Simply copy/paste that code and change the component name.

Refresh the page.

```tsx
/* eslint-disable */
if (process.env.NODE_ENV === 'development') {
  const whyDidYouRender = require('@welldone-software/why-did-you-render');
  whyDidYouRender(React, {
    logOnDifferentValues: true,
    trackAllPureComponents: true,
    trackHooks: true,
  });
}
COMPONENT.whyDidYouRender = true;
```
