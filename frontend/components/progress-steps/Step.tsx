import { cn } from '@/libs/utils';
import { toTitleCase } from '@/util/util';
import { CheckIcon } from '@radix-ui/react-icons';

interface Props {
  isCompleted: boolean;
  isActive: boolean;
  isLastStep: boolean;
  name: string;
}

export const Step = (props: Props) => {
  const { isActive, isCompleted, isLastStep, name } = props;

  return (
    <div className={cn(isLastStep ? 'flex' : 'flex-none')}>
      <StepCircle
        isCompleted={isCompleted}
        isLastStep={isLastStep}
        name={name}
        isActive={isActive}
      />
    </div>
  );
};

interface StepCircleProps {
  isCompleted: boolean;
  isLastStep: boolean;
  name: string;
  isActive: boolean;
}

const StepCircle = (props: StepCircleProps) => {
  const { isCompleted, isLastStep, name, isActive } = props;
  return (
    <div className="flex flex-row">
      <div className=" flex flex-col gap-2 items-center">
        <div
          className={cn(
            isActive || isCompleted
              ? 'bg-black '
              : 'border-2 border-black gray-800',
            'w-[20px] h-[20px]',
            'rounded-full',
            'justify-center flex align-middle items-center'
          )}
        >
          {isCompleted && <CheckIcon className="text-white" />}
        </div>
        <div className="text-xs w-[50px] justify-center flex">
          {toTitleCase(name)}
        </div>
      </div>
      {!isLastStep && (
        <div className=" w-[30px] h-[2px] mt-3 rounded-xl bg-gray-300" />
      )}
    </div>
  );
};