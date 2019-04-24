import { combineReducers } from 'redux';
import user from './user';
import modals from './modals';
import apps from './apps';
import app from './app';

const rootReducer = combineReducers({
  user,
  modals,
  apps,
  app
});

export default rootReducer;
