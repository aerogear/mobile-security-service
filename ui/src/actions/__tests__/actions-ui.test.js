import * as actions from '../actions-ui';
import {
  APPS_SORT,
  APP_VERSIONS_SORT,
  UPDATE_DISABLED_APP,
  UPDATE_VERSION_CUSTOM_MESSAGE,
  TOGGLE_NAVIGATION_MODAL,
  TOGGLE_APP_DETAILED_IS_DIRTY
} from '../types';
import { SortByDirection } from '@patternfly/react-table';

describe('actions', () => {
  it('should reverse the app Table Sort', () => {
    const index = 0;
    const direction = SortByDirection.asc;
    const expectedAction = {
      type: APPS_SORT,
      payload: {
        index: index,
        direction: direction
      }
    };
    expect(actions.appsTableSort(index, direction)).toEqual(expectedAction);
  });

  it('should reverse the app details Table Sort', () => {
    const index = 0;
    const direction = SortByDirection.asc;
    const expectedAction = {
      type: APP_VERSIONS_SORT,
      payload: {
        index: index,
        direction: direction
      }
    };
    expect(actions.appDetailsSort(index, direction)).toEqual(expectedAction);
  });

  it('should toggle the navigational modal', () => {
    const isOpen = true;
    const targetLocation = '/';
    const expectedAction = {
      type: TOGGLE_NAVIGATION_MODAL,
      payload: {
        isOpen: isOpen,
        targetLocation: targetLocation
      }
    };
    expect(actions.toggleNavigationModal(isOpen, targetLocation)).toEqual(expectedAction);
  });

  it('should toggle whether app detail view is dirty', () => {
    const expectedAction = {
      type: TOGGLE_APP_DETAILED_IS_DIRTY
    };
    expect(actions.toggleAppDetailedIsDirty()).toEqual(expectedAction);
  });

  it('should update the disabled status of the app version', () => {
    const id = 'v1.1';
    const isDisabled = false;
    const expectedAction = {
      type: UPDATE_DISABLED_APP,
      payload: {
        id: id,
        isDisabled: isDisabled
      }
    };
    expect(actions.updateDisabledAppVersion(id, isDisabled)).toEqual(expectedAction);
  });

  it('should update the custom text of the app version', () => {
    const id = 'v1.1';
    const value = false;
    const expectedAction = {
      type: UPDATE_VERSION_CUSTOM_MESSAGE,
      payload: {
        id: id,
        value: value
      }
    };
    expect(actions.updateVersionCustomMessage(id, value)).toEqual(expectedAction);
  });
});
