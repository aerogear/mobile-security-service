import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import { HeaderContainer } from './HeaderContainer';
import Header from '../components/Header';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      currentUser: {
        username: 'username',
        email: 'email'
      },
      isUserRequestFailed: false,
      getUser: jest.fn(),
      history: {
        push: jest.fn()
      }
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<HeaderContainer {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('HeaderContainer', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe(Header);
    });

    it('onTitleClick should call history push', () => {
      const { output, props } = setup();

      output.props.onTitleClick();
      expect(props.history.push).toBeCalled();
    });

    it('logout should call window.location.replace', () => {
      const { output } = setup();

      window.location.replace = jest.fn();
      output.props.logout();
      expect(window.location.replace).toBeCalledWith('/oauth/sign_in');
    });

    it('onUserDropdownToggle should change the isDropDownOpen state', () => {
      const { output } = setup();

      expect(output.props.isDropDownOpen).toBe(false);
    });
  });
});
