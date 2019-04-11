import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import LandingPage from './LandingPage';
import HeaderContainer from '../containers/HeaderContainer';
import AppsTableContainer from '../containers/AppsTableContainer';
import Content from './common/Content';
import { Title } from '@patternfly/react-core';

const setup = (propOverrides) => {
  const renderer = createRenderer();
  renderer.render(<LandingPage />);
  const output = renderer.getRenderOutput();
  return output;
};

describe('components', () => {
  describe('HeaderContainer', () => {
    it('should render', () => {
      const output = setup();
      const [ headerContainer ] = output.props.children;
      expect(headerContainer.type).toBe(HeaderContainer);
    });
  });

  describe('Content', () => {
    it('should render', () => {
      const output = setup();
      const [ , content ] = output.props.children;
      expect(content.type).toBe(Content);
    });
  });

  describe('Title', () => {
    it('should render', () => {
      const output = setup();
      const [ , content ] = output.props.children;
      const [ title ] = content.props.children;
      expect(title.type).toEqual(Title);
    });
  });

  describe('AppsTableContainer', () => {
    it('should render', () => {
      const output = setup();
      const [ , content ] = output.props.children;
      const [ , appsTableContainer ] = content.props.children;
      expect(appsTableContainer.type).toEqual(AppsTableContainer);
    });
  });
});
