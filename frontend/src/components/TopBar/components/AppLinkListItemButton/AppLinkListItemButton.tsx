import React from 'react';
import ListItemText from '@mui/material/ListItemText';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemButton from '@mui/material/ListItemButton';
import { useTranslation } from 'react-i18next';

export interface Props {
  link: string;
  primaryText: string;
  secondaryText?: string;
  iconElement?: React.ReactNode;
}

function AppLinkListItemButton({ link, primaryText, secondaryText = '', iconElement = null }: Props) {
  const { t } = useTranslation();

  return (
    <ListItemButton dense component="a" href={link}>
      {iconElement ? <ListItemIcon>{iconElement}</ListItemIcon> : null}
      <ListItemText primary={t(primaryText)} secondary={t(secondaryText)} />
    </ListItemButton>
  );
}

export default AppLinkListItemButton;
