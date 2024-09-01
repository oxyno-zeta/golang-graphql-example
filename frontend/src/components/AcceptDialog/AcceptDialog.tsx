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
  readonly open: boolean;
  readonly title: string;
  readonly content?: string;
  readonly contentElement?: ReactNode;
  readonly onClose: () => void;
  readonly onSubmit: () => Promise<void>;
  readonly okDisabled?: boolean;
  readonly dialogProps?: Partial<Omit<DialogProps, 'open' | 'onClose'>>;
  readonly dialogTitleProps?: Partial<DialogTitleProps>;
  readonly dialogContentProps?: Partial<DialogContentProps>;
  readonly dialogContentTextProps?: Partial<DialogContentTextProps>;
  readonly dialogActionsProps?: Partial<DialogActionsProps>;
  readonly cancelButtonProps?: Partial<Omit<ButtonProps, 'onClick'>>;
  readonly okButtonProps?: Partial<Omit<LoadingButtonProps, 'loading' | 'onClick' | 'disabled'>>;
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
  const [isLoading, setIsLoading] = useState<boolean>(false);
  // onClick ok
  const okOnClick = () => {
    setIsLoading(true);
    onSubmit().finally(() => {
      setIsLoading(false);
    });
  };

  return (
    <Dialog
      aria-describedby="alert-dialog-description"
      aria-labelledby="alert-dialog-title"
      onClose={onClose}
      open={open}
      {...dialogProps}
    >
      <DialogTitle id="alert-dialog-title" {...dialogTitleProps}>
        {title}
      </DialogTitle>
      <DialogContent id="alert-dialog-description" {...dialogContentProps}>
        {content ? <DialogContentText {...dialogContentTextProps}>{content}</DialogContentText> : null}
        {contentElement}
      </DialogContent>
      <DialogActions {...dialogActionsProps}>
        <Button onClick={onClose} {...cancelButtonProps}>
          {t('common.cancelAction')}
        </Button>
        <LoadingButton
          autoFocus
          disabled={okDisabled}
          loading={isLoading}
          onClick={okOnClick}
          variant="contained"
          {...okButtonProps}
        >
          {t('common.okAction')}
        </LoadingButton>
      </DialogActions>
    </Dialog>
  );
}

export default AcceptDialog;
