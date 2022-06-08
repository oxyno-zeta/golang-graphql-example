import React from 'react';
import { useTranslation } from 'react-i18next';
import Typography from '@mui/material/Typography';

function NotFoundRoute() {
  // Setup translate
  const { t } = useTranslation();

  return (
    <Typography sx={{ display: 'flex', justifyContent: 'center', textAlign: 'center', margin: '10px 0' }}>
      {t('common.routeNotFound')}
    </Typography>
  );
}

export default NotFoundRoute;
