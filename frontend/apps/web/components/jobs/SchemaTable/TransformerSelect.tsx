import { Button } from '@/components/ui/button';
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from '@/components/ui/command';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';
import { cn } from '@/libs/utils';
import {
  Transformer,
  isSystemTransformer,
  isUserDefinedTransformer,
} from '@/shared/transformers';
import { JobMappingTransformerForm } from '@/yup-validations/jobs';
import {
  SystemTransformer,
  UserDefinedTransformer,
  UserDefinedTransformerConfig,
} from '@neosync/sdk';
import { CaretSortIcon, CheckIcon } from '@radix-ui/react-icons';
import { ReactElement, useState } from 'react';

import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';

type Side = (typeof SIDE_OPTIONS)[number];

var SIDE_OPTIONS: readonly ['top', 'right', 'bottom', 'left'];

interface Props {
  transformers: Transformer[];
  value: JobMappingTransformerForm;
  onSelect(value: JobMappingTransformerForm): void;
  placeholder: string;
  side: Side;
  disabled: boolean;
}

export default function TransformerSelect(props: Props): ReactElement {
  const { transformers, value, onSelect, placeholder, side, disabled } = props;
  const [open, setOpen] = useState(false);

  const udfTransformers = transformers
    .filter(isUserDefinedTransformer)
    .sort((a, b) => a.name.localeCompare(b.name));
  const sysTransformers = transformers
    .filter(isSystemTransformer)
    .sort((a, b) => a.name.localeCompare(b.name));

  const udfTransformerMap = new Map(udfTransformers.map((t) => [t.id, t]));
  const sysTransformerMap = new Map(sysTransformers.map((t) => [t.source, t]));

  return disabled ? (
    <MockButtonWithToolTip
      placeholder={placeholder}
      value={value}
      udfTransformerMap={udfTransformerMap}
      systemTransformerMap={sysTransformerMap}
    />
  ) : (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          disabled={disabled}
          className={cn(
            placeholder.startsWith('Bulk')
              ? 'justify-between w-[275px]'
              : 'justify-between w-[175px]'
          )}
        >
          <div className="whitespace-nowrap truncate lg:w-[200px] text-left">
            {getPopoverTriggerButtonText(
              value,
              udfTransformerMap,
              sysTransformerMap,
              placeholder
            )}
          </div>
          <CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent
        className="w-[350px] p-0"
        avoidCollisions={false}
        side={side}
      >
        <Command>
          <CommandInput placeholder={placeholder} />
          <CommandEmpty>No transformers found.</CommandEmpty>
          <div className="max-h-[450px] overflow-y-scroll">
            <CommandGroup heading="Custom">
              {udfTransformers.map((t) => {
                return (
                  <CommandItem
                    key={t.id}
                    onSelect={() => {
                      onSelect({
                        source: 'custom',
                        config: {
                          case: 'userDefinedTransformerConfig',
                          value: new UserDefinedTransformerConfig({
                            id: t.id,
                          }),
                        },
                      });
                      setOpen(false);
                    }}
                    value={t.name}
                  >
                    <div className="flex flex-row items-center justify-between w-full">
                      <div className="flex flex-row items-center">
                        <CheckIcon
                          className={cn(
                            'mr-2 h-4 w-4',
                            value?.config?.case ===
                              'userDefinedTransformerConfig' &&
                              value?.source === 'custom' &&
                              value.config.value.id === t.id
                              ? 'opacity-100'
                              : 'opacity-0'
                          )}
                        />
                        <div className="items-center">{t?.name}</div>
                      </div>
                      <div className="ml-2 text-gray-400 text-xs">
                        {t.dataType}
                      </div>
                    </div>
                  </CommandItem>
                );
              })}
            </CommandGroup>
            <CommandGroup heading="System">
              {sysTransformers.map((t) => {
                return (
                  <CommandItem
                    key={t.source}
                    onSelect={() => {
                      let config = t.config?.config ?? {
                        case: '',
                        value: {},
                      };
                      if (!config.case) {
                        config = { case: '', value: {} };
                      }
                      onSelect({
                        source: t.source,
                        config: config,
                      });
                      setOpen(false);
                    }}
                    value={t.name}
                  >
                    <div className="flex flex-row items-center justify-between w-full">
                      <div className=" flex flex-row items-center">
                        <CheckIcon
                          className={cn(
                            'mr-2 h-4 w-4',
                            value?.source === t?.source
                              ? 'opacity-100'
                              : 'opacity-0'
                          )}
                        />
                        <div className="items-center">{t?.name}</div>
                      </div>
                      <div className="ml-2 text-gray-400 text-xs">
                        {t.dataType}
                      </div>
                    </div>
                  </CommandItem>
                );
              })}
            </CommandGroup>
          </div>
        </Command>
      </PopoverContent>
    </Popover>
  );
}

function getPopoverTriggerButtonText(
  value: JobMappingTransformerForm,
  udfTransformerMap: Map<string, UserDefinedTransformer>,
  systemTransformerMap: Map<string, SystemTransformer>,
  placeholder: string
): string {
  if (!value?.config) {
    return placeholder;
  }

  switch (value?.config?.case) {
    case 'userDefinedTransformerConfig':
      const id = value.config.value.id;
      return udfTransformerMap.get(id)?.name ?? placeholder;
    default:
      return systemTransformerMap.get(value.source)?.name ?? placeholder;
  }
}

interface MockButtonProps {
  value: JobMappingTransformerForm;
  placeholder: string;
  udfTransformerMap: Map<string, UserDefinedTransformer>;
  systemTransformerMap: Map<string, SystemTransformer>;
}

// have to do this otherwise react throws a warning where they don't want you to have a button as a child of a button (TooltipTrigger). If you do that and then disable the button the tooltip doesn't work on hover.
function MockButtonWithToolTip(props: MockButtonProps): ReactElement {
  const { placeholder, value, udfTransformerMap, systemTransformerMap } = props;

  return (
    <div>
      <TooltipProvider>
        <Tooltip>
          <TooltipTrigger asChild>
            <div
              className={cn(
                placeholder.startsWith('Bulk')
                  ? 'justify-between w-[275px]'
                  : 'justify-center w-[175px] cursor-pointer flex flex-row items-center  rounded-md text-sm font-gay- h-9 px-4 py-2 border border-input opacity-50'
              )}
            >
              <div className="whitespace-nowrap truncate lg:w-[200px] text-left">
                {getPopoverTriggerButtonText(
                  value,
                  udfTransformerMap,
                  systemTransformerMap,
                  placeholder
                )}
              </div>
              <CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50" />
            </div>
          </TooltipTrigger>
          <TooltipContent>
            <div className="max-w-[140px] text-wrap">
              Cannot assign a Transformer to Foreign Key if Primary Key has
              Transformer{' '}
            </div>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>
    </div>
  );
}
