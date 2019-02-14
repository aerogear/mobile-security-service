import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import './index.css';
import App from './components/App';
import configureStore from './configureStore';

const preloadedState = {
  apps: [
    { id: 1000, app: 'App-A', versions: 3, clients: 1245, startups: 'male', birth_year: '1979', actions: null },
    { id: 1001, app: 'App-B', versions: 4, clients: 655, startups: 'male', birth_year: '1974', actions: null },
    { id: 1002, app: 'App-C', versions: 1, clients: 970, startups: 'female', birth_year: '1989', actions: null },
    { id: 1003, app: 'App-D', versions: 6, clients: 255, startups: 'male', birth_year: '1990', actions: null },
    { id: 1004, app: 'App-E', versions: 5, clients: 120, startups: 'female', birth_year: '1999', actions: null }
  ]
};

const store = configureStore(preloadedState);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);
