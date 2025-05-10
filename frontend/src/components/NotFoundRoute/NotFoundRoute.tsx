import React from 'react';
import { useTranslation } from 'react-i18next';
import Typography, { TypographyProps } from '@mui/material/Typography';

export interface Props {
  readonly typographyProps?: TypographyProps;
}

function NotFoundRoute({ typographyProps = {} }: Props) {
  // Setup translate
  const { t } = useTranslation();

  return (
    <>
      <title>{t('common.routeNotFound')}</title>
      <Typography
        sx={{
          display: 'flex',
          justifyContent: 'center',
          textAlign: 'center',
          margin: '10px 0',
        }}
        {...typographyProps}
      >
        {t('common.routeNotFound')}
      </Typography>
    </>
  );
}

export default NotFoundRoute;
