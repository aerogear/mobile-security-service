export default class Version {
  constructor (json = {}) {
    this.id = json.id;
    this.versionNum = json.version;
    this.appId = json.appId;
    this.isDisabled = json.disabled;
    this.disabledMessage = json.disabledMessage;
    this.numOfAppLaunches = json.numOfAppLaunches;
    this.currentInstalls = json.numOfCurrentInstalls;
    this.lastLaunchedAt = json.lastLaunchedAt;
  }
}
