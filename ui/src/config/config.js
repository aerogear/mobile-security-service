// define a config object
const config = {
  mode: getEnv('NODE_ENV', 'development'),
  app: {
    name: getEnv('REACT_APP_NAME', 'Mobile Security Service')
  },
  dateTimeFormat: getEnv('REACT_APP_DATETIME_FORMAT', 'YYYY-MM-DD HH:mm:ss'),
  api: {
    url: getEnv('REACT_APP_API_URL', '/api')
  },
  auth: {
    url: getEnv('REACT_APP_AUTH_URL', '/oauth'),
    signInUrl: getEnv('REACT_APP_AUTH_SIGN_IN_URL', '/oauth/sign_in')
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
