import { combineReducers } from 'redux';
import header from './header';
import modals from './modals';
import apps from './apps';
import app from './app';

const rootReducer = combineReducers({
  header,
  modals,
  apps,
  app
});

export default rootReducer;
