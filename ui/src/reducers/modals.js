import {
  TOGGLE_NAVIGATION_MODAL,
  TOGGLE_SAVE_APP_MODAL,
  TOGGLE_DISABLE_APP_MODAL,
  DISABLE_APP_SUCCESS,
  SET_MODAL_DISABLE_MESSAGE
} from '../actions/types.js';

const initialState = {
  disableApp: {
    isOpen: false,
    disableMessage: ''
  },
  saveApp: {
    isOpen: false
  },
  navigationModal: {
    isOpen: false,
    targetLocation: undefined
  }
};

export default (state = initialState, action) => {
  switch (action.type) {
    case TOGGLE_NAVIGATION_MODAL: {
      const targetLocation = action.payload.targetLocation || undefined;
      return {
        ...state,
        navigationModal: {
          isOpen: action.payload.isOpen,
          targetLocation: targetLocation
        }
      };
    }
    case TOGGLE_SAVE_APP_MODAL: {
      return {
        ...state,
        saveApp: {
          isOpen: !state.saveApp.isOpen
        }
      };
    }
    case DISABLE_APP_SUCCESS: {
      return {
        ...state,
        disableApp: {
          isOpen: false
        }
      };
    }
    case TOGGLE_DISABLE_APP_MODAL: {
      return {
        ...state,
        disableApp: {
          isOpen: !state.disableApp.isOpen,
          disableMessage: ''
        }
      };
    }
    case SET_MODAL_DISABLE_MESSAGE: {
      return {
        ...state,
        disableApp: {
          ...state.disableApp,
          disableMessage: action.payload.disableMessage
        }
      };
    }
    default: {
      return state;
    }
  }
};
