import React from 'react';
import { connect } from 'react-redux';

import data from '../data';

const Index = () => {
  return (
    <div>
      <h1>LeetCode 150</h1>
      { data.map(item => (
        <div>
          <a href={ `/${item.id}` }>
            { `${item.id} ${item.title}` }
          </a>
        </div>
      )) }
    </div>
  );
};

const mapStateToProps = data => data;

export default connect(mapStateToProps)(Index);
