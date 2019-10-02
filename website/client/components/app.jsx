import React from 'react';
import styled from 'styled-components';
import { node } from 'prop-types';

import data from '../../data';

export const difficultyContainers = {
  1 : styled.span`
    background-color: #5cb85c;
  `,
  2 : styled.span`
    background-color: #f0ad4e;
  `,
  3 : styled.span`
    background-color: #d9534f;
  `
};
export const difficultyLabels = {
  1 : 'Easy',
  2 : 'Medium',
  3 : 'Hard'
};

const ListContainer = styled.div`
  display: block;
  font-size: 13px;
  height: 100vh;
  left: 0px;
  position: fixed;
  top: 0px;
  width: 350px;
  background: rgb(238, 238, 238);
  overflow: auto;
`;
const TitleContainer = styled.div`
  padding: 0 0 20px 0;
  text-align: center;

  & a {
    text-decoration: none;
    color: rgb(0, 173, 181);
  }
`;
const Link = styled.a`
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
  font-size: 13px;
  line-height: 2;
  color: inherit;
  display: block;
  border-top: 1px solid #ddd;
  text-decoration: none;

  :hover {
    background-color: #fff;
  }
`;
const LinkTextContainer = styled.div`
  padding: 10px 12px;
`;
const Content = styled.div`
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
  font-size: 18px;
  line-height: 1.5;
  box-sizing: inherit;
  display: block;
  margin-left: 350px;
  height: 100%;
`;

const App = ({ children }) => {
  return (
    <div>
      <ListContainer>
        <TitleContainer>
          <h1> LeetCode 150 </h1>
          <span>
          Language: &nbsp;
            <a
              target="_blank"
              href="https://golang.org/"
            >
          Go
            </a>
          </span>
        &nbsp;/&nbsp;
          <span>
          Author: &nbsp;
            <a
              target="_blank"
              href="https://xhu.me/"
            >
          xhu
            </a>
          </span>
        </TitleContainer>

        { data.map(item => {
          const DifficultyContainer = difficultyContainers[item.difficulty];
          const difficultyLabel = difficultyLabels[item.difficulty];

          return (
            <Link
              key={ item.id }
              href={ `/${item.id}` }
            >
              <LinkTextContainer>
                { `${item.id}. ${item.title}` }
                <DifficultyContainer
                  style={ {
                    display      : 'inline',
                    color        : '#fff',
                    fontSize     : 10,
                    padding      : '0 2px',
                    borderRadius : 3,
                    marginLeft   : 8
                  } }
                >
                  { difficultyLabel }
                </DifficultyContainer>
              </LinkTextContainer>
            </Link>
          );
        }) }
      </ListContainer>
      <Content>
        { children }
      </Content>
    </div>
  );
};

App.propTypes = {
  children : node.isRequired
};

export default App;
