import React from 'react';
import { Link as RouterLink } from 'react-router';
import Typography from '@mui/material/Typography';
import Link from '@mui/material/Link';
import { useTranslation } from 'react-i18next';
import type { BreadcrumbFixedDataConfig } from '../types';

interface Props {
  readonly breadcrumbData: BreadcrumbFixedDataConfig;
  readonly last: boolean;
  readonly pathname: string;
  readonly disablePageTitle?: boolean;
}

function FixedBreadcrumb({ breadcrumbData, last, pathname, disablePageTitle = false }: Props) {
  // Initialize translate
  const { t } = useTranslation();

  if (last) {
    return (
      <>
        {!disablePageTitle && <title>{t(breadcrumbData.textContent)}</title>}
        <Typography color="text.primary" {...(breadcrumbData.typographyProps || {})}>
          {t(breadcrumbData.textContent)}
        </Typography>
      </>
    );
  }

  return (
    <Link color="inherit" component={RouterLink} to={pathname} underline="hover" {...(breadcrumbData.linkProps || {})}>
      {t(breadcrumbData.textContent)}
    </Link>
  );
}

export default FixedBreadcrumb;
