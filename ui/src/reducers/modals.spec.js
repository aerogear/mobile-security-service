import modalsReducer from './modals';
import {
  TOGGLE_NAVIGATION_MODAL,
  TOGGLE_SAVE_APP_MODAL,
  TOGGLE_DISABLE_APP_MODAL,
  DISABLE_APP_SUCCESS
} from '../actions/types.js';

describe('modalsReducer', () => {
  const initialState = {
    disableApp: {
      isOpen: false
    },
    saveApp: {
      isOpen: false
    },
    navigationModal: {
      isOpen: false,
      targetLocation: undefined
    }
  };

  it('should return the initial state', () => {
    const newInitialState = JSON.stringify(modalsReducer(undefined, {}), null, 2);
    const expectedInitialState = JSON.stringify(initialState, null, 2);
    expect(newInitialState).toEqual(expectedInitialState);
  });

  it('should handle TOGGLE_NAVIGATION_MODAL open', () => {
    const isOpen = true;
    const targetLocation = '/';

    const newState = modalsReducer(initialState, {
      type: TOGGLE_NAVIGATION_MODAL,
      payload: {
        isOpen: isOpen,
        targetLocation: targetLocation
      }
    });
    expect(newState.navigationModal.isOpen).toEqual(isOpen);
    expect(newState.navigationModal.targetLocation).toEqual(targetLocation);
  });

  it('should handle TOGGLE_NAVIGATION_MODAL close', () => {
    const isOpen = true;

    const newState = modalsReducer(initialState, {
      type: TOGGLE_NAVIGATION_MODAL,
      payload: {
        isOpen: isOpen
      }
    });
    expect(newState.navigationModal.isOpen).toEqual(isOpen);
    expect(newState.navigationModal.targetLocation).toEqual(undefined);
  });

  it('should handle TOGGLE_SAVE_APP_MODAL', () => {
    const isOpen = initialState.saveApp.isOpen;
    const newState = modalsReducer(initialState, { type: TOGGLE_SAVE_APP_MODAL });
    expect(newState.saveApp.isOpen).toEqual(!isOpen);
  });

  it('should handle DISABLE_APP_SUCCESS', () => {
    const newState = modalsReducer(initialState, {
      type: DISABLE_APP_SUCCESS
    });
    expect(newState.disableApp.isOpen).toEqual(false);
  });

  it('should handle TOGGLE_DISABLE_APP_MODAL', () => {
    const isOpen = initialState.disableApp.isOpen;
    const newState = modalsReducer(initialState, { type: TOGGLE_DISABLE_APP_MODAL });
    expect(newState.disableApp.isOpen).toEqual(!isOpen);
  });
});
