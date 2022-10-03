import React, { ReactNode } from 'react';
import Divider from '@mui/material/Divider';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';
import { TopBarSpacer } from '../../TopBar';

interface Props {
  titleText?: string;
  titleElement?: string;
  subtitleElement?: ReactNode;
  listItemsElement: ReactNode;
}

const defaultProps = {
  titleText: '',
  titleElement: undefined,
  subtitleElement: undefined,
};

function DrawerContent({ titleText, titleElement, subtitleElement, listItemsElement }: Props) {
  return (
    <div style={{ paddingTop: '10px' }}>
      <TopBarSpacer />
      <div style={{ margin: '10px' }}>
        <div
          style={{
            display: 'flex',
            alignItems: 'center',
            textAlign: 'center',
            justifyContent: 'center',
          }}
        >
          {titleText && (
            <Typography
              style={{
                fontWeight: 'bold',
                textOverflow: 'ellipsis',
                whiteSpace: 'nowrap',
                width: '100%',
                overflow: 'hidden',
              }}
            >
              {titleText}
            </Typography>
          )}
          {titleElement}
        </div>

        {subtitleElement && (
          <div
            style={{
              marginTop: '5px',
              display: 'flex',
              alignItems: 'center',
              textAlign: 'center',
              justifyContent: 'center',
            }}
          >
            {subtitleElement}
          </div>
        )}
      </div>

      <div style={{ margin: '10px 25px 0 25px' }}>
        <Divider />
      </div>

      <List>{listItemsElement}</List>
    </div>
  );
}

DrawerContent.defaultProps = defaultProps;

export default DrawerContent;
