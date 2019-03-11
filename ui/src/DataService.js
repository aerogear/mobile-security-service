const baseUrl = '/api';

export const wsError = {};

const requestConfig = (method, body) => ({
  method,
  cache: 'no-cache',
  credentials: 'same-origin',
  headers: {
    'Content-Type': 'application/json; charset=utf-8'
  },
  body: body && JSON.stringify(body)
});

const request = async (url, method, body) => {
  const response = await fetch(`${baseUrl}/${url}`, requestConfig(method, body));
  if (!response.ok) {
    const msg = await response.text();
    throw Error(msg);
  }
  return method === 'DELETE' || response.json();
};

const fetchItems = async url => {
  const result = await request(url, 'GET');
  return result || [];
};

const fetchItem = async url => {
  const result = await request(url, 'GET');
  return result || {};
};

const dataService = {
  fetchApps: () => fetchItems('apps'),
  getAppById: (id) => fetchItem(`apps/${id}`)
};

export default dataService;
