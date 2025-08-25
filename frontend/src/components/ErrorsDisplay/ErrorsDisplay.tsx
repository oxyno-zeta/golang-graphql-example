import React, { type Key, type ReactNode } from 'react';
import Typography, { type TypographyProps } from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { ServerError, CombinedGraphQLErrors } from '@apollo/client';
import { useTranslation } from 'react-i18next';
import type { SxProps } from '@mui/material';
import { GraphqlErrorsExtensionsCodeCustomComponentMapKeyPrefix, ServerErrorCustomComponentMapKey } from './constants';

export interface Props {
  readonly error?: Error | null;
  readonly errors?: Error[];
  readonly noMargin?: boolean;
  readonly containerBoxSx?: SxProps;
  readonly errorTitleTypographyProps?: TypographyProps;
  readonly errorElementTypographyProps?: TypographyProps;
  readonly ulSx?: SxProps;
  readonly liSx?: SxProps;
  readonly customErrorComponents?: Record<string, React.ElementType>;
  readonly customErrorComponentProps?: Record<string, object>;
}

export interface CustomNetworkError {
  message: string;
  path: null | undefined | readonly (string | number)[];
}

function ErrorsDisplay({
  error = null,
  errors = [],
  noMargin = false,
  containerBoxSx = {},
  errorTitleTypographyProps = {},
  errorElementTypographyProps = {},
  ulSx = {},
  liSx = {},
  customErrorComponents = {},
  customErrorComponentProps = {},
}: Props) {
  // Initialize translate
  const { t } = useTranslation();

  // Check if error or errors are set
  if ((!error || !errors) && errors && errors.length === 0) {
    // Display nothing
    return null;
  }

  // Build error list
  let errorList = errors;
  // Check if error is set
  if (error) {
    // Override array
    errorList = [error];
  }

  const elements: ReactNode[] = [];

  // Loop over errors
  errorList?.forEach((err, mainIndex) => {
    const contents: { content: string; key: Key; CustomComp?: React.ElementType; customCompProps?: object }[] = [];
    if (ServerError.is(err) || CombinedGraphQLErrors.is(err)) {
      if (ServerError.is(err)) {
        contents.push({
          content: err.message,
          key: mainIndex,
          CustomComp: customErrorComponents[ServerErrorCustomComponentMapKey],
          customCompProps: customErrorComponentProps[ServerErrorCustomComponentMapKey],
        });
      }

      if (CombinedGraphQLErrors.is(err)) {
        contents.push(
          ...err.errors.map(({ message, extensions }, i) => {
            let mess = message;
            let customComp: React.ElementType | undefined;
            let customCompProps: object | undefined;
            // Check if there is a code in extensions
            if (extensions && extensions.code) {
              mess = t(`common.errorCode.${extensions.code}`);

              const mapKey = `${GraphqlErrorsExtensionsCodeCustomComponentMapKeyPrefix}${extensions.code}`;

              customComp = customErrorComponents[mapKey];
              customCompProps = customErrorComponentProps[mapKey];
            }

            return {
              content: mess,
              key: `${mainIndex}-graphQLErrors-${i}`,
              CustomComp: customComp,
              customCompProps,
            };
          }),
        );
      }
    } else {
      // Isn't an apollo error
      contents.push({ content: err.message, key: mainIndex });
    }

    if (contents) {
      // Check if a content have been detected
      // Save element
      elements.push(
        ...contents.map(({ content, key, CustomComp, customCompProps }) => (
          <Box component="li" key={key} sx={liSx}>
            {CustomComp ? (
              <CustomComp {...customCompProps} />
            ) : (
              <Typography color="error" {...errorElementTypographyProps}>
                {content}
              </Typography>
            )}
          </Box>
        )),
      );
    }
  });

  return (
    <Box sx={{ color: 'error', margin: noMargin ? 0 : 8, ...containerBoxSx }}>
      <Typography color="error" {...errorTitleTypographyProps}>
        {t('common.errors')}:
      </Typography>
      <Box component="ul" sx={ulSx}>
        {elements}
      </Box>
    </Box>
  );
}

export default ErrorsDisplay;
