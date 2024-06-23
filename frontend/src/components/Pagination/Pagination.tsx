import React, { ForwardRefExoticComponent, RefAttributes } from 'react';
import { Link, useHref, useLocation, createSearchParams, useSearchParams, LinkProps } from 'react-router-dom';
import { mdiChevronLeft, mdiChevronRight, mdiChevronDoubleLeft } from '@mdi/js';
import SvgIcon from '@mui/material/SvgIcon';
import Tooltip from '@mui/material/Tooltip';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import Toolbar, { ToolbarProps } from '@mui/material/Toolbar';
import { useTranslation } from 'react-i18next';
import { PageInfoModel } from '../../models/general';
import { getAllSearchParams } from '../../utils/urlSearchParams';
import { cleanPaginationSearchParams } from '../../utils/pagination';

export interface Props {
  pageInfo: PageInfoModel;
  maxPaginationSize: number;
  // Using onFirstPage will disable search param management
  onFirstPage?: () => void | undefined;
  // Using onPreviousPage will disable search param management
  onPreviousPage?: () => void | undefined;
  // Using onNextPage will disable search param management
  onNextPage?: () => void | undefined;
  toolbarProps?: ToolbarProps;
  firstIconButtonProps?: IconButtonProps;
  previousIconButtonProps?: IconButtonProps;
  nextIconButtonProps?: IconButtonProps;
}

type IconButtonInternalProps = {
  to?: string;
  component?: ForwardRefExoticComponent<LinkProps & RefAttributes<HTMLAnchorElement>>;
  onClick?: () => void;
};

function Pagination({
  maxPaginationSize,
  pageInfo,
  onFirstPage = undefined,
  onPreviousPage = undefined,
  onNextPage = undefined,
  toolbarProps = {},
  firstIconButtonProps = {},
  nextIconButtonProps = {},
  previousIconButtonProps = {},
}: Props) {
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
  if (onFirstPage) {
    firstPageProps = {
      onClick: onFirstPage,
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
  if (onPreviousPage) {
    previousPageProps = {
      onClick: onPreviousPage,
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
  if (onNextPage) {
    nextPageProps = {
      onClick: onNextPage,
    };
  }

  return (
    <Toolbar variant="dense" {...toolbarProps}>
      <div style={{ marginLeft: 'auto' }}>
        <Tooltip title={<>{t('common.firstPageAction')}</>}>
          <span>
            <IconButton disabled={!pageInfo.hasPreviousPage} {...firstPageProps} {...firstIconButtonProps}>
              <SvgIcon>
                <path d={mdiChevronDoubleLeft} />
              </SvgIcon>
            </IconButton>
          </span>
        </Tooltip>
        <Tooltip title={<>{t('common.previousPageAction')}</>}>
          <span>
            <IconButton disabled={!pageInfo.hasPreviousPage} {...previousPageProps} {...previousIconButtonProps}>
              <SvgIcon>
                <path d={mdiChevronLeft} />
              </SvgIcon>
            </IconButton>
          </span>
        </Tooltip>
        <Tooltip title={<>{t('common.nextPageAction')}</>}>
          <span>
            <IconButton disabled={!pageInfo.hasNextPage} {...nextPageProps} {...nextIconButtonProps}>
              <SvgIcon>
                <path d={mdiChevronRight} />
              </SvgIcon>
            </IconButton>
          </span>
        </Tooltip>
      </div>
    </Toolbar>
  );
}

export default Pagination;
