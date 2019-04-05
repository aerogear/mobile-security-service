import * as types from './types';
import * as actions from './actions-ui';

describe('actions', () => {
  it('appsTableSort should create APPS_SORT action', () => {
    const index = 0;
    const direction = 'ASC';
    expect(actions.appsTableSort(null, index, direction)).toEqual({
      type: types.APPS_SORT,
      payload: {
        index: index,
        direction: direction
      }
    });
  });

  it('appVersionsTableSort should create APP_VERSIONS_SORT action', () => {
    const index = 1;
    const direction = 'DESC';
    expect(actions.appVersionsTableSort(null, index, direction)).toEqual({
      type: types.APP_VERSIONS_SORT,
      payload: {
        index: index,
        direction: direction
      }
    });
  });

  it('setAppDetailedDirtyState should create SET_APP_DETAILED_DIRTY action', () => {
    const isDirty = true;
    expect(actions.setAppDetailedDirtyState(isDirty)).toEqual({
      type: types.SET_APP_DETAILED_DIRTY,
      payload: {
        isDirty: isDirty
      }
    });
  });

  it('updateVersionCustomMessage should create UPDATE_VERSION_CUSTOM_MESSAGE action', () => {
    const id = '123';
    const value = 'customMessage';
    expect(actions.updateVersionCustomMessage(id, value)).toEqual({
      type: types.UPDATE_VERSION_CUSTOM_MESSAGE,
      payload: {
        id: id,
        value: value
      }
    });
  });

  it('updateDisabledAppVersion should create UPDATE_DISABLED_APP action', () => {
    const id = '123';
    const isDisabled = true;
    expect(actions.updateDisabledAppVersion(id, isDisabled)).toEqual({
      type: types.UPDATE_DISABLED_APP,
      payload: {
        id: id,
        isDisabled: isDisabled
      }
    });
  });

  it('toggleNavigationModal should create TOGGLE_NAVIGATION_MODAL action', () => {
    const isOpen = true;
    const targetLocation = '/';
    expect(actions.toggleNavigationModal(isOpen, targetLocation)).toEqual({
      type: types.TOGGLE_NAVIGATION_MODAL,
      payload: {
        isOpen: isOpen,
        targetLocation: targetLocation
      }
    });
  });

  it('toggleSaveAppModal should create TOGGLE_SAVE_APP_MODAL action', () => {
    expect(actions.toggleSaveAppModal()).toEqual({
      type: types.TOGGLE_SAVE_APP_MODAL
    });
  });

  it('toggleDisableAppModal should create TOGGLE_DISABLE_APP_MODAL action', () => {
    expect(actions.toggleDisableAppModal()).toEqual({
      type: types.TOGGLE_DISABLE_APP_MODAL
    });
  });
});
