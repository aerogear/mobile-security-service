import { createStore, applyMiddleware } from 'redux';
import thunkMiddleware from 'redux-thunk';
import { createLogger } from 'redux-logger';
import appsReducer from './reducers/apps';

let middleware = [thunkMiddleware];
if (process.env.NODE_ENV !== 'production') {
  middleware = [...middleware, createLogger()];
}

export default function configureStore (preloadedState) {
  return createStore(appsReducer, preloadedState, applyMiddleware(...middleware));
}
