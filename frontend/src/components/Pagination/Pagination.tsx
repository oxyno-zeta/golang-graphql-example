import React from 'react';
import ChevronRight from '@mui/icons-material/ChevronRight';
import ChevronLeft from '@mui/icons-material/ChevronLeft';
import Tooltip from '@mui/material/Tooltip';
import IconButton from '@mui/material/IconButton';
import Toolbar from '@mui/material/Toolbar';
import { useTranslation } from 'react-i18next';
import { PageInfoModel, PaginationInputModel } from '../../models/general';

interface Props {
  pageInfo: PageInfoModel;
  maxPaginationSize: number;
  setPaginationInput: (v: PaginationInputModel) => void;
}

function Pagination({ maxPaginationSize, pageInfo, setPaginationInput }: Props) {
  // Setup translate
  const { t } = useTranslation();

  const handlePreviousPage = () => {
    setPaginationInput({
      before: pageInfo.startCursor,
      last: maxPaginationSize,
      first: undefined,
      after: undefined,
    });
  };

  const handleNextPage = () => {
    setPaginationInput({
      before: undefined,
      last: undefined,
      first: maxPaginationSize,
      after: pageInfo.endCursor,
    });
  };

  return (
    <Toolbar variant="dense">
      <div style={{ marginLeft: 'auto' }}>
        <Tooltip title={<>{t('common.previousPageAction')}</>}>
          <span>
            <IconButton disabled={!pageInfo.hasPreviousPage} onClick={handlePreviousPage}>
              <ChevronLeft />
            </IconButton>
          </span>
        </Tooltip>
        <Tooltip title={<>{t('common.nextPageAction')}</>}>
          <span>
            <IconButton disabled={!pageInfo.hasNextPage} onClick={handleNextPage}>
              <ChevronRight />
            </IconButton>
          </span>
        </Tooltip>
      </div>
    </Toolbar>
  );
}

export default Pagination;
