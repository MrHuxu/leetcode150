import 'core-js/stable';

import React from 'react';
import { renderToString } from 'react-dom/server';

import { Provider } from 'react-redux';
import { createStaticStore } from './store';

import { Switch, Route, StaticRouter } from 'react-router';
import routes from './routes';

import { ServerStyleSheet, StyleSheetManager } from 'styled-components';

import App from './components/app';
const sheet = new ServerStyleSheet();

export const renderHtmlString = (url, data) => {
  const html = renderToString(
    <StyleSheetManager sheet={ sheet.instance }>
      <Provider store={ createStaticStore(data) }>
        <StaticRouter location={ url }>
          <App>
            <Switch>
              { routes.map(route => (
                <Route { ...route } />
              )) }
            </Switch>
          </App>
        </StaticRouter>
      </Provider>
    </StyleSheetManager>
  );
  const styles = sheet.getStyleTags();
  return styles + html;
};
