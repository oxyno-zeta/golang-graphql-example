import React, { type ReactNode } from 'react';
import Divider from '@mui/material/Divider';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';

export interface Props {
  readonly titleText?: string;
  readonly titleElement?: string;
  readonly subtitleElement?: ReactNode;
  readonly listItemsElement: ReactNode;
  readonly isNormalCollapsed: boolean;
}

function DrawerContent({
  titleText = '',
  titleElement = undefined,
  subtitleElement = undefined,
  listItemsElement,
  isNormalCollapsed,
}: Props) {
  return (
    <>
      {isNormalCollapsed ? (
        <>
          <div style={{ margin: '10px' }}>
            <div
              style={{
                display: 'flex',
                alignItems: 'center',
                textAlign: 'center',
                justifyContent: 'center',
              }}
            >
              {titleText ? (
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
              ) : null}
              {titleElement}
            </div>

            {subtitleElement ? (
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
            ) : null}
          </div>

          <div style={{ margin: '0 25px' }}>
            <Divider />
          </div>
        </>
      ) : null}

      <List>{listItemsElement}</List>
    </>
  );
}

export default DrawerContent;
