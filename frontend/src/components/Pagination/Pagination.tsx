import React, { ForwardRefExoticComponent, RefAttributes } from 'react';
import { Link, useHref, useLocation, createSearchParams, useSearchParams, LinkProps } from 'react-router-dom';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import KeyboardDoubleArrowLeftIcon from '@mui/icons-material/KeyboardDoubleArrowLeft';
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
  handleFirstPage?: () => void | undefined;
  handlePreviousPage?: () => void | undefined;
  handleNextPage?: () => void | undefined;
}

const defaultProps = {
  handleFirstPage: undefined,
  handlePreviousPage: undefined,
  handleNextPage: undefined,
};

type IconButtonInternalProps = {
  to?: string;
  component?: ForwardRefExoticComponent<LinkProps & RefAttributes<HTMLAnchorElement>>;
  onClick?: () => void;
};

function Pagination({ maxPaginationSize, pageInfo, handleFirstPage, handlePreviousPage, handleNextPage }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // Get search params
  const [searchParams] = useSearchParams();
  // Setup location
  const location = useLocation();

  // Build first page props
  let firstPageProps: IconButtonInternalProps = {
    to: useHref({
      pathname: location.pathname,
      search: createSearchParams(getAllSearchParams(cleanPaginationSearchParams(searchParams))).toString(),
    }),
    component: Link,
  };
  // Check if handle first page is declared
  if (handleFirstPage) {
    firstPageProps = {
      onClick: handleFirstPage,
    };
  }

  // Build previous page props
  let previousPageProps: IconButtonInternalProps = {
    to: useHref({
      pathname: location.pathname,
      search: createSearchParams({
        ...getAllSearchParams(cleanPaginationSearchParams(searchParams)),
        before: pageInfo.startCursor || '',
        last: maxPaginationSize.toString(),
      }).toString(),
    }),
    component: Link,
  };
  // Check if handle previous page is declared
  if (handlePreviousPage) {
    previousPageProps = {
      onClick: handlePreviousPage,
    };
  }

  // Build next page props
  let nextPageProps: IconButtonInternalProps = {
    to: useHref({
      pathname: location.pathname,
      search: createSearchParams({
        ...getAllSearchParams(cleanPaginationSearchParams(searchParams)),
        after: pageInfo.endCursor || '',
        first: maxPaginationSize.toString(),
      }).toString(),
    }),
    component: Link,
  };
  // Check if handle next page is declared
  if (handleNextPage) {
    nextPageProps = {
      onClick: handleNextPage,
    };
  }

  return (
    <Toolbar variant="dense">
      <div style={{ marginLeft: 'auto' }}>
        <Tooltip title={<>{t('common.firstPageAction')}</>}>
          <span>
            <IconButton disabled={!pageInfo.hasPreviousPage} {...firstPageProps}>
              <KeyboardDoubleArrowLeftIcon />
            </IconButton>
          </span>
        </Tooltip>
        <Tooltip title={<>{t('common.previousPageAction')}</>}>
          <span>
            <IconButton disabled={!pageInfo.hasPreviousPage} {...previousPageProps}>
              <ChevronLeftIcon />
            </IconButton>
          </span>
        </Tooltip>
        <Tooltip title={<>{t('common.nextPageAction')}</>}>
          <span>
            <IconButton disabled={!pageInfo.hasNextPage} {...nextPageProps}>
              <ChevronRightIcon />
            </IconButton>
          </span>
        </Tooltip>
      </div>
    </Toolbar>
  );
}

Pagination.defaultProps = defaultProps;

export default Pagination;
