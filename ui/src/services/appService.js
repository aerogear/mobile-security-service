
/**
 * Gets the difference between two arrays of app versions.
 *
 * @param {Array} savedVersions - The original app version data
 * @param {*} currentVersions - The versions which are editable
 */
const getDirtyVersions = (savedVersions, currentVersions) => {
  return currentVersions.filter((v, i) => {
    const currentVersionValues = Object.values(v).toString();
    const savedVersionValues = Object.values(savedVersions[i]).toString();

    return currentVersionValues !== savedVersionValues;
  });
};

const AppService = {
  getDirtyVersions: (savedVersions, currentVersions) => getDirtyVersions(savedVersions, currentVersions)
};

export default AppService;
