import React, { ReactNode, useState } from 'react';
import Button, { ButtonProps } from '@mui/material/Button';
import Dialog, { DialogProps } from '@mui/material/Dialog';
import LoadingButton, { LoadingButtonProps } from '@mui/lab/LoadingButton';
import DialogActions, { DialogActionsProps } from '@mui/material/DialogActions';
import DialogContent, { DialogContentProps } from '@mui/material/DialogContent';
import DialogContentText, { DialogContentTextProps } from '@mui/material/DialogContentText';
import DialogTitle, { DialogTitleProps } from '@mui/material/DialogTitle';
import { useTranslation } from 'react-i18next';

export interface Props {
  open: boolean;
  title: string;
  content?: string;
  contentElement?: ReactNode;
  onClose: () => void;
  onSubmit: () => Promise<void>;
  okDisabled?: boolean;
  dialogProps?: Partial<Omit<DialogProps, 'open' | 'onClose'>>;
  dialogTitleProps?: Partial<DialogTitleProps>;
  dialogContentProps?: Partial<DialogContentProps>;
  dialogContentTextProps?: Partial<DialogContentTextProps>;
  dialogActionsProps?: Partial<DialogActionsProps>;
  cancelButtonProps?: Partial<Omit<ButtonProps, 'onClick'>>;
  okButtonProps?: Partial<Omit<LoadingButtonProps, 'loading' | 'onClick' | 'disabled'>>;
}

function AcceptDialog({
  open,
  title,
  content = '',
  contentElement = null,
  onClose,
  onSubmit,
  okDisabled = false,
  dialogProps = {},
  dialogTitleProps = {},
  dialogContentProps = {},
  dialogContentTextProps = {},
  dialogActionsProps = {},
  cancelButtonProps = {},
  okButtonProps = {},
}: Props) {
  const { t } = useTranslation();
  // Manage loading
  const [isLoading, setLoading] = useState<boolean>(false);
  // onClick ok
  const okOnClick = () => {
    setLoading(true);
    onSubmit().finally(() => {
      setLoading(false);
    });
  };

  return (
    <Dialog
      open={open}
      onClose={onClose}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
      {...dialogProps}
    >
      <DialogTitle id="alert-dialog-title" {...dialogTitleProps}>
        {title}
      </DialogTitle>
      <DialogContent id="alert-dialog-description" {...dialogContentProps}>
        {content && <DialogContentText {...dialogContentTextProps}>{content}</DialogContentText>}
        {contentElement}
      </DialogContent>
      <DialogActions {...dialogActionsProps}>
        <Button onClick={onClose} {...cancelButtonProps}>
          {t('common.cancelAction')}
        </Button>
        <LoadingButton
          loading={isLoading}
          variant="contained"
          onClick={okOnClick}
          disabled={okDisabled}
          autoFocus
          {...okButtonProps}
        >
          {t('common.okAction')}
        </LoadingButton>
      </DialogActions>
    </Dialog>
  );
}

export default AcceptDialog;
