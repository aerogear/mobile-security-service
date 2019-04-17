import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import { AppPageContainer } from './AppPageContainer';
import AppOverview from '../../components/appView/AppOverview';
import Content from '../../components/common/Content';
import AppToolbar from '../../components/appView/AppToolbar';
import HeaderContainer from '../HeaderContainer';
import AppVersionsTableContainer from './AppVersionsTableContainer';
import DisableAppModalContainer from './DisableAppModalContainer';
import NavigationModalContainer from './NavigationModalContainer';
import SaveAppModalContainer from './SaveAppModalContainer';
import { Title } from '@patternfly/react-core';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      match: {
        params: {
          id: 1
        }
      },
      app: {
        appId: 'com.test',
        deployedVersions: []
      },
      savedData: {
        deployedVersions: []
      },
      isDirty: false,
      getAppById: jest.fn(),
      history: {
        block: jest.fn()
      },
      toggleNavigationModal: jest.fn(),
      toggleSaveAppModal: jest.fn(),
      toggleDisableAppModal: jest.fn(),
      saveAppVersions: jest.fn(),
      setAppDetailedDirtyState: jest.fn()
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<AppPageContainer {...props} />);
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
      const [ headerContainer ] = output.props.children;
      expect(headerContainer.type).toBe(HeaderContainer);
    });
  });

  describe('AppToolbar', () => {
    it('should render', () => {
      const { output } = setup();
      const [ , , appToolbar ] = output.props.children;
      expect(appToolbar.type).toBe(AppToolbar);
    });
  });

  describe('Content', () => {
    it('should render', () => {
      const { output } = setup();
      const [ , , , content ] = output.props.children;
      expect(content.type).toBe(Content);
    });

    describe('AppOverview', () => {
      it('should render', () => {
        const { output } = setup();
        const [ , , , content ] = output.props.children;
        const [ appOverview ] = content.props.children;
        expect(appOverview.type).toBe(AppOverview);
      });
    });

    describe('Title', () => {
      it('should render', () => {
        const { output } = setup();
        const [ , , , content ] = output.props.children;
        const [ , title ] = content.props.children;
        expect(title.type).toBe(Title);
      });
    });

    describe('AppVersionsTableContainer', () => {
      it('should render', () => {
        const { output } = setup();
        const [ , , , content ] = output.props.children;
        const [ , , appVersionsTableContainer ] = content.props.children;
        expect(appVersionsTableContainer.type).toBe(AppVersionsTableContainer);
      });
    });

    describe('NavigationModalContainer', () => {
      it('should render', () => {
        const { output } = setup();
        const [ , , , content ] = output.props.children;
        const [ , , , navigationModalContainer ] = content.props.children;
        expect(navigationModalContainer.type).toBe(NavigationModalContainer);
      });
    });

    describe('SaveAppModalContainer', () => {
      it('should render', () => {
        const { output } = setup();
        const [ , , , content ] = output.props.children;
        const [ , , , , saveAppModalContainer ] = content.props.children;
        expect(saveAppModalContainer.type).toBe(SaveAppModalContainer);
      });
    });

    describe('DisableAppModalContainer', () => {
      it('should render', () => {
        const { output } = setup();
        const [ , , , content ] = output.props.children;
        const [ , , , , , disableAppModalContainer ] = content.props.children;
        expect(disableAppModalContainer.type).toBe(DisableAppModalContainer);
      });
    });
  });
});

describe('events', () => {
  describe('AppToolbar handles clicks', () => {
    it('onSaveAppClick should call toggleSaveAppModal', () => {
      const { output, props } = setup();
      const [ , , appToolbar ] = output.props.children;
      appToolbar.props.onSaveAppClick();
      expect(props.toggleSaveAppModal).toBeCalled();
    });

    it('onDisableAppClick should call toggleDisableAppModal', () => {
      const { output, props } = setup();
      const [ , , appToolbar ] = output.props.children;
      appToolbar.props.onDisableAppClick();
      expect(props.toggleDisableAppModal).toBeCalled();
    });
  });
});
