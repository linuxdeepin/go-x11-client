# Run tests in check section
# disable for bootstrapping
%bcond_with check
%global with_debug 1

%if 0%{?with_debug}
%global debug_package   %{nil}
%endif

%global provider        github
%global provider_tld    com
%global project         linuxdeepin
%global repo            go-x11-client

%global provider_prefix %{provider}.%{provider_tld}/%{project}/%{repo}
%global import_path     %{provider_prefix}

Name:           golang-github-linuxdeepin-go-x11-client
Version:        0.6.7
Release:        1
Summary:        A X11 client Go bindings for Deepin Desktop Environment
License:        GPLv3
URL:            %{gourl}
Source0:        %{name}_%{version}.orig.tar.xz	
BuildRequires:  compiler(go-compiler)

%description
%{summary}.

%package devel
Summary:        %{summary}
BuildArch:      noarch
BuildRequires:  gocode


%description devel
%{summary}.


%prep
%forgeautosetup

%install
install -d -p %{buildroot}/%{gopath}/src/%{import_path}/
for file in $(find . -iname "*.go") ; do
    install -d -p %{buildroot}/%{gopath}/src/%{import_path}/$(dirname $file)
    cp -pav $file %{buildroot}/%{gopath}/src/%{import_path}/$file
    echo "%%{gopath}/src/%%{import_path}/$file" >> devel.file-list
done

%files devel -f devel.file-list
%doc README
%license LICENSE

%changelog
* Thu Mar 19 2021 uoser <uoser@uniontech.com> - 0.6.7-1
- Update to 0.6.7
