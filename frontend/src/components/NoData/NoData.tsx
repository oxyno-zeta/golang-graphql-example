import React from 'react';
import { useTranslation } from 'react-i18next';
import Typography, { TypographyProps } from '@mui/material/Typography';

export interface Props {
  typographyProps?: Omit<TypographyProps, 'children'>;
}

function NoData({ typographyProps = {} }: Props) {
  // Setup translate
  const { t } = useTranslation();

  return (
    <Typography
      sx={{ display: 'flex', justifyContent: 'center', textAlign: 'center', margin: '10px 0' }}
      {...typographyProps}
    >
      {t('common.noData')}
    </Typography>
  );
}

export default NoData;
