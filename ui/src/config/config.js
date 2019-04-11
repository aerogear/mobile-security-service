// define a config object
const config = {
  app: {
    name: getEnv('REACT_APP_NAME', 'Mobile Security Service')
  },
  dateTimeFormat: getEnv('REACT_APP_DATETIME_FORMAT', 'YYYY-MM-DD HH:mm:ss'),
  api: {
    url: getEnv('REACT_APP_API_URL', '/api')
  }
};

/**
 * Return an environment variable or use a default value
 *
 * @param {string} key is the environment variable key
 * @param {string} defaultVal is the default value to fall back on
 * @returns {string} the config option value
 */
function getEnv (key, defaultVal) {
  const value = process.env[key];

  return value || defaultVal;
}

export default config;
