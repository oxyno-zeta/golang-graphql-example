import React from 'react';
import { Link, useHref, useLocation, createSearchParams, useSearchParams } from 'react-router-dom';
import ChevronRight from '@mui/icons-material/ChevronRight';
import ChevronLeft from '@mui/icons-material/ChevronLeft';
import Tooltip from '@mui/material/Tooltip';
import IconButton from '@mui/material/IconButton';
import Toolbar from '@mui/material/Toolbar';
import { useTranslation } from 'react-i18next';
import { PageInfoModel } from '../../models/general';
import { getAllSearchParams } from '../../utils/urlSearchParams';
import { cleanPaginationSearchParams } from '../../utils/pagination';

interface Props {
  pageInfo: PageInfoModel;
  maxPaginationSize: number;
}

function Pagination({ maxPaginationSize, pageInfo }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // Get search params
  const [searchParams] = useSearchParams();
  // Setup location
  const location = useLocation();

  return (
    <Toolbar variant="dense">
      <div style={{ marginLeft: 'auto' }}>
        <Tooltip title={<>{t('common.previousPageAction')}</>}>
          <span>
            <IconButton
              component={Link}
              disabled={!pageInfo.hasPreviousPage}
              to={useHref({
                pathname: location.pathname,
                search: createSearchParams({
                  ...getAllSearchParams(cleanPaginationSearchParams(searchParams)),
                  before: pageInfo.startCursor || '',
                  last: maxPaginationSize.toString(),
                }).toString(),
              })}
            >
              <ChevronLeft />
            </IconButton>
          </span>
        </Tooltip>
        <Tooltip title={<>{t('common.nextPageAction')}</>}>
          <span>
            <IconButton
              component={Link}
              disabled={!pageInfo.hasNextPage}
              to={useHref({
                pathname: location.pathname,
                search: createSearchParams({
                  ...getAllSearchParams(cleanPaginationSearchParams(searchParams)),
                  after: pageInfo.endCursor || '',
                  first: maxPaginationSize.toString(),
                }).toString(),
              })}
            >
              <ChevronRight />
            </IconButton>
          </span>
        </Tooltip>
      </div>
    </Toolbar>
  );
}

export default Pagination;
