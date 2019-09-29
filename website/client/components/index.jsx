import React from 'react';

import data from '../../data';

const Index = () => {
  console.log({ data });
  return (
    <div>
      <h1>LeetCode 150</h1>
      { data.map(item => (
        <div>
          <a href={ `https://leetcode.com/problems/${item.slug}/` } target="_blank">
            { `${item.id} ${item.title}` }
          </a>
        </div>
      )) }
    </div>
  );
};

export default Index;
