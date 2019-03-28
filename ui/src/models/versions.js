export default class Version {
  constructor (json = {}) {
    this.id = json.id;
    this.version = json.version;
    this.appId = json.appId;
    this.disabled = json.disabled;
    this.disabledMessage = json.disabledMessage;
    this.numOfAppLaunches = json.numOfAppLaunches;
    this.currentInstalls = json.numOfCurrentInstalls;
    this.lastLaunchedAt = json.lastLaunchedAt;
  }

  getId () {
    console.log('getting id: ', this.id);
    return this.id;
  }

  isDisabled () {
    return this.disabled;
  }

  getDisabledMessage () {
    return this.disabledMessage;
  }

  getInstalls () {
    return this.currentInstalls;
  }
  getVersion () {
    return this.version;
  }

  getAppLaunches () {
    return this.numOfAppLaunches;
  }

  getLastLaunchedAt () {
    return this.lastLaunchedAt;
  }

  setIsDisabled (disabled) {
    this.disabled = disabled;
  }

  setDisabledMessage (message) {
    this.disabledMessage = message;
  }
}
