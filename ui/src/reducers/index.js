import { GET_APPS } from '../actions/types';

const intialState = { apps: [] };

export default function apps (state = intialState, action) {
  switch (action.type) {
    case GET_APPS:
      return {
        apps: action.payload
      };
    default:
      return state;
  }
}
