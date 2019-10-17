import React from 'react';

import data from '../../data';

const mapDiffToLabel = {
  1 : 'Easy',
  2 : 'Medium',
  3 : 'Hard'
};

const App = () => {
  return (
    <div>
      <div style={ { padding: '0 0 15px 0' } }>
        <span><h1> LeetCode 150 </h1></span>
        <span>
          Language: &nbsp;
          <a target="_blank" href="https://golang.org/" >
          Go
          </a>
        </span>
        &nbsp;/&nbsp;
        <span>
          Author: &nbsp;
          <a target="_blank" href="https://xhu.me/" >
          xhu
          </a>
        </span>
      </div>

      { data.map(item => {
        return (
          <div className="link-container">
            <a key={ item.id } href={ `/${item.id}` }>
              { `${item.id}. ${item.title}` }
              <span className={ `difficulty difficulty-${item.difficulty}` }>
                { mapDiffToLabel[item.difficulty] }
              </span>
            </a>
          </div>
        );
      }) }
    </div>
  );
};

export default App;
