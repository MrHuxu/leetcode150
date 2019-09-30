import React from 'react';

import data from '../data';

const Index = () => {
  return (
    <div>
      <h1>LeetCode 150</h1>
      { data.map(item => (
        <div>
          <a
            key={ item.id }
            href={ `/${item.id}` }
          >
            { `${item.id} ${item.title}` }
          </a>
        </div>
      )) }
    </div>
  );
};

export default Index;
