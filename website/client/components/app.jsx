import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Route, Redirect, Link } from 'react-router-dom';
import { push } from 'react-router-redux';
import { Style } from 'radium';
import { Button } from 'antd';

import { shape, bool, object, string, func } from 'prop-types';

import Get from './get';
import Post from './post';
import Random from './random';
import Index from './index';

import { global, btnArea } from '../styles/app';

const mapDispatchToProps = { push };

/**
 * Use stateful component for enabling the hot module reload
 * HMR is not working for stateless component and it will reload the whole page
 */

@connect(null, mapDispatchToProps)
class App extends Component {
  render () {
    const { match, push } = this.props;
    return (
      <div>
        <Style rules={ global } />
        <Index />
      </div>
    );
  }
}

App.propTypes = {
  match : shape({
    isExact : bool.isRequired,
    params  : object,
    path    : string.isRequired,
    url     : string.isRequired
  }).isRequired,
  push : func.isRequired
};

export default App;
