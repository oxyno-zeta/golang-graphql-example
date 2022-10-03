import React, { ReactNode, useState } from 'react';
import Button from '@mui/material/Button';
import Dialog, { DialogProps } from '@mui/material/Dialog';
import LoadingButton from '@mui/lab/LoadingButton';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { useTranslation } from 'react-i18next';

interface Props {
  open: boolean;
  title: string;
  content?: string;
  contentElement?: ReactNode;
  handleClose: () => void;
  handleOk: () => Promise<void>;
  okDisabled?: boolean;
  dialogProps?: Partial<DialogProps>;
}

const defaultProps = {
  content: '',
  contentElement: null,
  okDisabled: false,
  dialogProps: {},
};

function AcceptDialog({ open, title, content, contentElement, handleClose, handleOk, okDisabled, dialogProps }: Props) {
  const { t } = useTranslation();
  // Manage loading
  const [isLoading, setLoading] = useState<boolean>(false);
  // onClick ok
  const okOnClick = () => {
    setLoading(true);
    handleOk().finally(() => {
      setLoading(false);
    });
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
      {...dialogProps}
    >
      <DialogTitle id="alert-dialog-title">{title}</DialogTitle>
      <DialogContent id="alert-dialog-description">
        {content && <DialogContentText>{content}</DialogContentText>}
        {contentElement}
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>{t('common.cancelAction')}</Button>
        <LoadingButton loading={isLoading} variant="contained" onClick={okOnClick} disabled={okDisabled} autoFocus>
          {t('common.okAction')}
        </LoadingButton>
      </DialogActions>
    </Dialog>
  );
}
// Add default props
AcceptDialog.defaultProps = defaultProps;

export default AcceptDialog;
