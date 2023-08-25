import { MouseEvent, useEffect, useRef } from "react";
import styled from "styled-components";
import styles from '@/components/styles/DialogModal.module.css'

const Buttons = styled.div`
  display: flex;
  gap: 20px;
`;

const isClickInsideRectangle = (e: MouseEvent, element: HTMLElement) => {
  const r = element.getBoundingClientRect();

  return (
    e.clientX > r.left &&
    e.clientX < r.right &&
    e.clientY > r.top &&
    e.clientY < r.bottom
  );
};

type Props = {
  title: string;
  isOpened: boolean;
  onProceed: () => void;
  onClose: () => void;
  children: React.ReactNode;
  proceedButtonLabel?: string;
  cancelButtonLabel?: string;
};

const DialogModal = ({
  title,
  isOpened,
  onProceed,
  onClose,
  children,
  proceedButtonLabel,
  cancelButtonLabel,
}: Props) => {
  const ref = useRef<HTMLDialogElement>(null);

  useEffect(() => {
    if (isOpened) {
      ref.current?.showModal();
      document.body.classList.add("modal-open"); // prevent bg scroll
    } else {
      ref.current?.close();
      document.body.classList.remove("modal-open");
    }
  }, [isOpened]);

  const proceedAndClose = () => {
    onProceed();
    onClose();
  };

  return (
  <div className = {styles.dialog_container}>
    <dialog className = {styles.dialog}
      ref={ref}
      onCancel={onClose}
      onClick={(e) =>
        ref.current && !isClickInsideRectangle(e, ref.current) && onClose()
      }
    >
      <h3>{title}</h3>

      {children}

      <Buttons>
        <center>
          <button className={styles.dialog_button} onClick={proceedAndClose}>{proceedButtonLabel || "Ok"}</button>
          <button className={styles.dialog_button} onClick={onClose}>{cancelButtonLabel || "Cancel"}</button>
        </center>
      </Buttons>
    </dialog>
  </div>
  );
};

export default DialogModal;
