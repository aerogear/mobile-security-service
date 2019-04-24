import config from './config/config';
import { requestConfig } from './DataService';

/**
 * Sends a HTTP request
 *
 * @param {string} path - The URI to send the request to
 * @param {*} method - HTTP Method (PUT, GET, PATCH, DELETE, POST)
 * @param {*} body - The request payload
 * @returns {boolean|JSON} - Returns the response from the HTTP request
 */
const request = async (path, method, body) => {
  const response = await fetch(`${config.auth.url}/${path}`, requestConfig(method, body));
  if (!response.ok) {
    const msg = await response.text();

    throw Error(msg);
  }

  return response;
};

/**
 * Makes a HTTP GET request to a provided URI and returns a single item
 *
 * @param {string} path - The path to make the GET request to
 */
const fetchItem = async path => {
  const result = await request(path, 'GET');
  return result || {};
};

const AuthService = {
  authenticate: () => fetchItem('auth')
};

export default AuthService;
