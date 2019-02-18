import { createStore, applyMiddleware } from 'redux';
import thunkMiddleware from 'redux-thunk';
import { createLogger } from 'redux-logger';
import appsReducer from './reducers/index.js';

let middleware = [thunkMiddleware];
if (process.env.NODE_ENV !== 'production') {
  middleware = [...middleware, createLogger()];
}

export default function configureStore () {
  return createStore(appsReducer, applyMiddleware(...middleware));
}
