import React from 'react';
import { useTranslation } from 'react-i18next';
import Typography, { TypographyProps } from '@mui/material/Typography';

export interface Props {
  typographyProps?: TypographyProps;
}

const defaultProps = {
  typographyProps: {},
};

function NotFoundRoute({ typographyProps }: Props) {
  // Setup translate
  const { t } = useTranslation();

  return (
    <Typography
      sx={{ display: 'flex', justifyContent: 'center', textAlign: 'center', margin: '10px 0' }}
      {...typographyProps}
    >
      {t('common.routeNotFound')}
    </Typography>
  );
}

NotFoundRoute.defaultProps = defaultProps;

export default NotFoundRoute;
