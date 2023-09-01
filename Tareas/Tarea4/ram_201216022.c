#include <linux/init.h>
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/sysinfo.h>
#include <linux/mm.h>


MODULE_LICENSE("GPL");
MODULE_AUTHOR("Tu Nombre");
MODULE_DESCRIPTION("Módulo de información de RAM");
MODULE_VERSION("0.1");

static int carnet = 12345;

static int ram_info_proc_show(struct seq_file *m, void *v) {
    struct sysinfo info;
    si_meminfo(&info);

    seq_printf(m, "Total de RAM: %lu KB\n", info.totalram * (info.mem_unit / 1024));
    seq_printf(m, "Memoria RAM en uso: %lu KB\n", (info.totalram - info.freeram) * (info.mem_unit / 1024));
    seq_printf(m, "Memoria RAM libre: %lu KB\n", info.freeram * (info.mem_unit / 1024));
    seq_printf(m, "Porcentaje de uso: %lu%%\n", ((info.totalram - info.freeram) * 100) / info.totalram);

    return 0;
}

static int ram_info_proc_open(struct inode *inode, struct file *file) {
    return single_open(file, ram_info_proc_show, NULL);
}

static const struct proc_ops ram_info_proc_fops = {
    .open = ram_info_proc_open,
    .read = seq_read,
    .llseek = seq_lseek,
    .release = single_release,
    .proc_write = NULL,
};

static int __init ram_module_init(void) {
    printk(KERN_INFO "Número de carnet: %d\n", carnet);
    proc_create("ram_info", 0, NULL, &ram_info_proc_fops);
    return 0;
}

static void __exit ram_module_exit(void) {
    remove_proc_entry("ram_info", NULL);
    printk(KERN_INFO "Nombre del estudiante: Tu Nombre\n");
}

module_init(ram_module_init);
module_exit(ram_module_exit);
