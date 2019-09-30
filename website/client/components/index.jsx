import React from 'react';
import styled from 'styled-components';

import data from '../data';

const Container = styled.div`
  background-color: #eee;
`;
const Left = styled.div`
  position: fixed;
  height: 100%;
  width: 30%;
  text-align: center;
  padding: 24px 0 0 0;
  color: #555;
`;
const Right = styled.div`
  padding: 24px 0 30px 30%;
`;
const LinkContainer = styled.div`
  margin: 12px 0;
`;
const Link = styled.a`
  color: #666;
  font-size: 15px;
`;
const difficultyContainers = {
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
const difficultyLabels = {
  1 : 'Easy',
  2 : 'Medium',
  3 : 'Hard'
};

const Index = () => {
  return (
    <Container>
      <Left>
        <h1>LeetCode 150</h1>
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
                  fontSize     : 11,
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
