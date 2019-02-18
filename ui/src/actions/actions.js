import { GET_APPS } from './types';

export const getApps = () => dispatch => {
  const apps = [
    [ 'App-F', 3, 245, 4 ],
    [ 'App-G', 4, 655, 5 ],
    [ 'App-H', 1, 970, 6 ],
    [ 'App-I', 6, 255, 7 ],
    [ 'App-J', 5, 120, 8 ]
  ];
  // We could do a fetch here on the backend to get the updated state of apps
  dispatch({
    type: GET_APPS,
    payload: apps
  });
};
