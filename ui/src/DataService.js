import config from './config/config';

/**
 * Creates a configuration obect with information about the request type
 * and the request body
 *
 * @param {string} method - HTTP Method (PUT, GET, PATCH, DELETE, POST)
 * @param {object|Array} body - JSON Payload to send in the request
 * @returns a config object to send in HTTP requests
 */
const requestConfig = (method, body) => ({
  method,
  cache: 'no-cache',
  credentials: 'same-origin',
  headers: {
    'Content-Type': 'application/json; charset=utf-8'
  },
  body: body && JSON.stringify(body)
});

/**
 * Sends a HTTP request
 *
 * @param {string} path - The URI to send the request to
 * @param {*} method - HTTP Method (PUT, GET, PATCH, DELETE, POST)
 * @param {*} body - The request payload
 * @returns {boolean|JSON} - Returns the response from the HTTP request
 */
const request = async (path, method, body) => {
  const response = await fetch(`${config.api.url}/${path}`, requestConfig(method, body));
  if (!response.ok) {
    const msg = await response.text();
    throw Error(`${response.statusText}: ${msg}`);
  }

  return (method === 'DELETE' || response.status === 204) || response.json();
};

/**
 * Makes a HTTP GET request to a provided URI and returns an array of items
 *
 * @param {string} url - The full URL to make the GET request to
 */
const fetchItems = async url => {
  const result = await request(url, 'GET');
  return result || [];
};

/**
 * Makes a HTTP GET request to a provided URI and returns a single item
 *
 * @param {string} url - The full URL to make the GET request to
 */
const fetchItem = async url => {
  const result = await request(url, 'GET');
  return result || {};
};

/**
 * Makes a PUT request to update a collection of items.
 *
 * @param {string} url - The path to the REST endpoint
 * @param {Array} body - The array of items to update
 */
const putItems = async (url, body) => {
  const result = await request(url, 'PUT', body);
  return result || [];
};

/**
 * Disable all app versions in the server. Optionally set a
 * custom disabled message for all app versions.
 *
 * @param {string} id - The App ID to disable all versions for
 * @param {string} [disabledMessage] - Sets the disabled message for all versions to this.
 */
const disableAppVersions = async (id, disabledMessage) => {
  await request(`apps/${id}/versions/disable`, 'POST', { disabledMessage });

  return disabledMessage;
};

const dataService = {
  fetchApps: () => fetchItems('apps'),
  getAppById: (id) => fetchItem(`apps/${id}`),
  updateAppVersions: (id, versions) => putItems(`apps/${id}/versions`, versions),
  disableAppVersions
};

export default dataService;
