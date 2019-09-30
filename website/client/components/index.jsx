import React from 'react';
import styled from 'styled-components';

import data from '../../data';
import { difficultyContainers, difficultyLabels } from './common';

const Container = styled.div`
  background-color: #eee;
`;
const Left = styled.div`
  position: fixed;
  height: 100%;
  width: 30%;
  text-align: center;
  padding: 45px 0 0 0;
  color: #555;
`;
const Right = styled.div`
  padding: 24px 0 30px 30%;
`;
const LinkContainer = styled.div`
  margin: 1.15rem 0;
`;
const Link = styled.a`
  color: #666;
`;

const Index = () => {
  return (
    <Container>
      <Left>
        <h1>LeetCode 150</h1>
        <div>
          Language: &nbsp;
          <a
            target="_blank"
            href="https://golang.org/"
          >
          Go
          </a>
        </div>
        <div>
          Author: &nbsp;
          <a
            target="_blank"
            href="https://xhu.me/"
          >
          xhu
          </a>
        </div>
      </Left>

      <Right>
        { data.map(item => {
          const DifficultyContainer = difficultyContainers[item.difficulty];
          const difficultyLabel = difficultyLabels[item.difficulty];

          return (
            <LinkContainer>
              <Link
                key={ item.id }
                href={ `/${item.id}` }
              >
                { `${item.id}. ${item.title}` }
              </Link>
              <DifficultyContainer
                style={ {
                  color        : '#fff',
                  fontSize     : '0.8rem',
                  padding      : '0 2px',
                  borderRadius : 3,
                  marginLeft   : 8
                } }
              >
                { difficultyLabel }
              </DifficultyContainer>
            </LinkContainer>
          );
        }) }
      </Right>
    </Container>
  );
};

export default Index;
