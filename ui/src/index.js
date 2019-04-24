import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import '@patternfly/react-core/dist/styles/base.css';

import './index.css';
import App from './containers/App';
import configureStore from './configureStore';

// load environment variables from dotenv
require('dotenv').config();

const preloadedState = {};

const store = configureStore(preloadedState);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);
