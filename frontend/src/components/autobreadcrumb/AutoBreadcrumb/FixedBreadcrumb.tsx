import React from 'react';
import { type Params, Link as RouterLink } from 'react-router';
import Typography from '@mui/material/Typography';
import Link from '@mui/material/Link';
import { useTranslation } from 'react-i18next';
import type { BreadcrumbFixedDataConfig } from '../types';

interface Props {
  readonly breadcrumbData: BreadcrumbFixedDataConfig;
  readonly last: boolean;
  readonly pathname: string;
  readonly params: Params<string>;
  readonly disablePageTitle?: boolean;
}

function FixedBreadcrumb({ breadcrumbData, last, pathname, params, disablePageTitle = false }: Props) {
  // Initialize translate
  const { t } = useTranslation();

  const computedPath = breadcrumbData.overrideComputedPath
    ? breadcrumbData.overrideComputedPath(pathname, params)
    : pathname;

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
    <Link
      color="inherit"
      component={RouterLink}
      to={computedPath}
      underline="hover"
      {...(breadcrumbData.linkProps || {})}
    >
      {t(breadcrumbData.textContent)}
    </Link>
  );
}

export default FixedBreadcrumb;
